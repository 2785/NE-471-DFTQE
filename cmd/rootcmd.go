package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/2785/n471-proj-carrot/internal/nanohub"
	"github.com/2785/n471-proj-carrot/internal/strain"
	"github.com/2785/n471-proj-carrot/model"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RootCmd(cmd *cobra.Command, args []string) error {

	// 1. Load Config
	fmt.Printf("Loading Config...")
	if !(viper.IsSet("inputdir") && viper.IsSet("outputdir") && viper.IsSet("numberOfAxis") && viper.IsSet("numberOfStressVariation")) {
		return errors.New("input or output not set")
	}

	fmt.Println("  done")

	// 2. Load nanohub input files

	inputdir := viper.GetString("inputdir")

	files, err := ioutil.ReadDir(inputdir)
	if err != nil {
		return fmt.Errorf("cannot read input directory: %w", err)
	}

	fmt.Printf("discovered %v files\n", len(files))

	simulations := make([]*model.Simulation, len(files))

	fmt.Printf("Start Loading Input Files...")

	for i, f := range files {
		b, err := ioutil.ReadFile(path.Join(inputdir, f.Name()))
		if err != nil {
			return fmt.Errorf("cannot read input file %s: %w", f, err)
		}

		input, err := strain.Decode(string(b))

		if err != nil {
			return fmt.Errorf("cannot decode input file %s: %w", f, err)
		}
		simulations[i] = &model.Simulation{
			Input: *input,
		}
	}

	fmt.Println("  done")

	// 3. Sort output files

	type outputFile struct {
		Name       string
		Start, End int
	}

	outputdir := viper.GetString("outputdir")

	files, err = ioutil.ReadDir(outputdir)

	if err != nil {
		return fmt.Errorf("cannot read output directory: %w", err)
	}

	fmt.Printf("discovered %v files\n", len(files))

	fmt.Printf("parsing file names assuming format 'test_N_to_M_Band.txt' and 'test_N_to_M_DOS.txt'\n")

	re := regexp.MustCompile(`test_(?P<start>\d+)_to_(?P<end>\d+)_(?P<type>[^\.]+).txt`)

	dosSrc := []outputFile{}
	bandSrc := []outputFile{}

	for _, f := range files {
		match := re.FindStringSubmatch(f.Name())
		if match == nil {
			return errors.New("no match")
		}

		result := make(map[string]string)

		for i, name := range re.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}

		for _, v := range []string{"start", "end", "type"} {
			val, ok := result[v]
			if !ok || val == "" {
				return fmt.Errorf("%s not found", v)
			}
		}

		start, err := strconv.Atoi(strings.TrimSpace(result["start"]))
		if err != nil {
			return fmt.Errorf("could not parse start value: %w", err)
		}
		end, err := strconv.Atoi(strings.TrimSpace(result["end"]))
		if err != nil {
			return fmt.Errorf("could not parse end value: %w", err)
		}

		switch result["type"] {
		case "DOS":
			dosSrc = append(dosSrc, outputFile{f.Name(), start, end})
		case "Band":
			bandSrc = append(bandSrc, outputFile{f.Name(), start, end})
		default:
			return fmt.Errorf("unknown file type %s in file %s", result["type"], f.Name())
		}
	}

	fmt.Println("... done")

	fmt.Printf("found %v dos files and %v band files\n", len(dosSrc), len(bandSrc))

	// 4. Load DOS Files

	dosSplitter := `------------------------------------------------------------
 Density of States
------------------------------------------------------------`

	fmt.Printf("loading DoS files...")
	for _, f := range dosSrc {
		start := f.Start
		rev := f.End < f.Start
		count := func() int {
			if rev {
				start = f.End
				return f.Start - f.End + 1
			}
			return f.End - f.Start + 1
		}()
		b, err := ioutil.ReadFile(path.Join(outputdir, f.Name))
		if err != nil {
			return fmt.Errorf("error reading file '%s'", f.Name)
		}
		parts := strings.Split(string(b), dosSplitter)[1:]
		if len(parts) != count {
			return fmt.Errorf("unexpected number of parts in file '%s', expected %v, found %v", f.Name, count, len(parts))
		}

		if rev {
			reverseAny(parts)
		}

		for i, p := range parts {
			dosInfo, err := nanohub.DecodeDoS(p)
			if err != nil {
				return fmt.Errorf("error decoding dos file %s section %v: %w", f.Name, i+1, err)
			}
			simulations[start-1+i].DoS = *dosInfo
		}
	}
	fmt.Println("  done")

	// 5. Load Band Structure Files

	bandSplitter := regexp.MustCompile(`--+\n Band 1\n--+\nK-P.*\n`)

	fmt.Printf("Loading Band Files...")

	for _, f := range bandSrc {
		start := f.Start
		rev := f.End < f.Start
		count := func() int {
			if rev {
				start = f.End
				return f.Start - f.End + 1
			}
			return f.End - f.Start + 1
		}()
		b, err := ioutil.ReadFile(path.Join(outputdir, f.Name))
		if err != nil {
			return fmt.Errorf("error reading file '%s'", f.Name)
		}
		parts := bandSplitter.Split(string(b), -1)

		if len(parts) == 1 {
			return fmt.Errorf("error splitting band file %s", f.Name)
		}

		parts = parts[1:]

		if len(parts) != count {
			return fmt.Errorf("unexpected number of parts in file '%s', expected %v, found %v", f.Name, count, len(parts))
		}

		if rev {
			reverseAny(parts)
		}

		for i, p := range parts {
			bandInfo, err := nanohub.DecodeBands(p)
			if err != nil {
				return fmt.Errorf("error decoding dos file %s section %v: %w", f.Name, i+1, err)
			}
			simulations[start-1+i].Bands = *bandInfo
		}
	}

	fmt.Println("  done")

	numAxes := viper.GetInt("numberOfAxis")
	numSets := viper.GetInt("numberOfStressVariation")

	thingySets := func() [][]*model.Simulation {
		out := make([][]*model.Simulation, numSets)
		for i := 0; i < numSets; i++ {
			out[i] = simulations[i*numAxes : (i+1)*numAxes]
		}
		return out
	}()

	someOtherThingySet := simulations[numAxes*numSets:]

	_, _, _, _ = numAxes, numSets, thingySets, someOtherThingySet

	_ = simulations

	return nil
}

func reverseAny(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
