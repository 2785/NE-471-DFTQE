from io import StringIO
from re import error
import pandas as pd
import re
import matplotlib.pyplot as plt
import seaborn as sns

sns.set_style(style="darkgrid")

runs = dict()

fermiLevels = dict()

f = open("../testdata/DensityOfStates.txt", "r")

key = 0

while (l := f.readline()) != "":
    if l.startswith("---"):
        continue

    if l.find("Density of States") != -1:
        f.readline()
        match = re.findall("\d", f.readline())
        if len(match) == 0:
            raise error("nothing found")
        ind = int(match[0])
        key = ind

        runs[key] = list()
        continue

    if l.find("Fermi level") != -1:
        f.readline()
        f.readline()
        fermiLevels[key] = float(f.readline().split(",")[0].strip())
        f.readline()

    if l.find(",") != -1:
        runs[key].append(l)

dfs = dict()

for k, v in runs.items():
    dfs[k] = pd.read_csv(StringIO("".join(v)), names=["energy", "dos"])

print(dfs.get(1).head())
print(fermiLevels)

p1 = sns.relplot(x="energy", y="dos", data=dfs.get(1), kind="line", markers=True)

sns.relplot(x="energy", y="dos", data=dfs.get(2), kind="line", markers=True)


p1.set_titles
p1.set_axis_labels(x_var="Energy (eV)", y_var="Density of States")

plt.show()