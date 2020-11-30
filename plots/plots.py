import pandas as pd
import matplotlib.pyplot as plt

### DOS ALL ###

data = pd.read_csv(
    "./csv_out/dosAll.csv",
    sep=",",
    dtype={"energy": float},
)

plt.style.use("dark_background")

# Plot Uniform Stress on the principle Axes vs non-uniform 1 v 3
plt.plot(data["Energy (eV)"], data[f"Run #{1}"], label="Uniform Stress")
plt.plot(data["Energy (eV)"], data[f"Run #{49}"], label="Non-uniform Stress")
plt.title("DOS for Uniform Stress vs Non Uniform Stress on Principle Axes")
plt.xlabel("Energy (ev)")
plt.ylabel("Density of States")
plt.legend()
plt.savefig("dos_plot1v3.png")
plt.close()

# Plot uniform stress vs uniform stress non-uniform sheer 1 v 2
plt.plot(data["Energy (eV)"], data[f"Run #{1}"], label="Uniform Stress")
plt.plot(data["Energy (eV)"], data[f"Run #{25}"], label="Uniform Stress Plus Sheer")
plt.title("DOS for Uniform Stress vs Uniform Stress and Non-uniform Sheer")
plt.xlabel("Energy (ev)")
plt.ylabel("Density of States")
plt.legend()
plt.savefig("dos_plot1v2.png")
plt.close()

# Plot uniform stress and Sheer vs non-uniform stress non-uniform sheer 3 v 4
plt.plot(data["Energy (eV)"], data[f"Run #{49}"], label="Uniform Stress Plus Sheer")
plt.plot(data["Energy (eV)"], data[f"Run #{73}"], label="Non-Uniform Stress Plus Sheer")
plt.title("DOS for Non-Uniform Stress vs Uniform Stress and Sheer On Principle Axes")
plt.xlabel("Energy (ev)")
plt.ylabel("Density of States")
plt.legend()
plt.savefig("dos_plot3v4.png")
plt.close()

# Plot uniform stress on principle axes vs stress and sheer 1 v 5
plt.plot(data["Energy (eV)"], data[f"Run #{1}"], label="Uniform Stress")
plt.plot(data["Energy (eV)"], data[f"Run #{97}"], label="Stress and Sheer")
plt.title("DOS for Uniform Stress on Principle Axes vs Stress and Sheer")
plt.xlabel("Energy (ev)")
plt.ylabel("Density of States")
plt.legend()
plt.savefig("dos_plot1v5.png")
plt.close()


### Fermi ALL ###

data = pd.read_csv(
    "./csv_out/fermiAll.csv",
    sep=",",
    dtype={"energy": float},
)

plt.plot(
    data["Run"],
    data["Energy (eV)"],
    color="r",
    marker="o",
    markersize=10,
    linestyle="None",
    label='Reference Axes on [0,0,0]'
)

plt.title("Fermi Level for Different Stress Strain Combinations")
plt.xlabel("Run Number")
plt.ylabel("Fermi Energy (eV)")
plt.legend()
plt.savefig("fermi_plot.png")
plt.close()


### BandGap ALL ###

data = pd.read_csv(
    "./csv_out/bandGapAll.csv",
    sep=",",
    dtype={"energy": float},
)

plt.plot(
    data["Run"],
    data["Band Gap Energy (eV)"],
    color="r",
    marker="o",
    markersize=10,
    linestyle="None",
    label='Reference Axes on [0,0,0]'
)

plt.title("Fermi Level for Different Stress Strain Combinations")
plt.xlabel("Run Number")
plt.ylabel("Band Gap Energy (eV)")
plt.legend()
plt.savefig("bandGap_plot.png")
plt.close()

### DOS Increasing Stress ###

data = pd.read_csv(
    "./csv_out/dos.csv",
    sep=",",
    dtype={"energy": float},
)

# Plot Increasing Stress on the principle Axes
plt.plot(data["Energy (eV)"], data[f"Run #{121}"], label="0 %")
plt.plot(data["Energy (eV)"], data[f"Run #{124}"], label="33%")
plt.plot(data["Energy (eV)"], data[f"Run #{127}"], label="66%")
plt.plot(data["Energy (eV)"], data[f"Run #{130}"], label="100%")
plt.title("DOS for Increasing Stress Along Principle Axes (4GPA)")
plt.xlabel("Energy (ev)")
plt.ylabel("Density of States")
plt.legend()
plt.savefig("dos_inc.png")
plt.close()

### Fermi Inc ###

data = pd.read_csv(
    "./csv_out/fermi.csv",
    sep=",",
    dtype={"energy": float},
)

# update run dataset to use correct percentage of max
for i in range(len(data["Run"])):
    data["Run"][i] = ((data["Run"][i] - 1) * 4/9) / 4

plt.plot(
    data["Run"],
    data["Energy (eV)"],
    label='Reference Axes on [0,0,0]'
)

plt.title("Fermi Level for Increasing Uniform Stress (Max 4 GPA)")
plt.xlabel("Percentage of Max Stress")
plt.ylabel("Fermi Energy (eV)")
plt.legend()
plt.savefig("fermi_inc.png")
plt.close()


### BandGap Inc ###

data = pd.read_csv(
    "./csv_out/bandGap.csv",
    sep=",",
    dtype={"energy": float},
)

# update run dataset to use correct percentage of max
for i in range(len(data["Run"])):
    data["Run"][i] = ((data["Run"][i] - 121) * 4/9) / 4

plt.plot(
    data["Run"],
    data["Band Gap Energy (eV)"],
    label='Reference Axes on [0,0,0]'
)

plt.title("Band Gap for Increasing Uniform Stress (Max 4 GPA)")
plt.xlabel("Percentage of Max Stress")
plt.ylabel("Band Gap Energy (eV)")
plt.legend()
plt.savefig("bandGap_inc.png")
plt.close()