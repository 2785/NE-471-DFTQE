import pandas as pd
import matplotlib.pyplot as plt
import copy

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

# Plot uniform stress vs uniform stress non-uniform Shear 1 v 2
plt.plot(data["Energy (eV)"], data[f"Run #{1}"], label="Uniform Stress")
plt.plot(data["Energy (eV)"], data[f"Run #{25}"], label="Uniform Stress Plus Shear")
plt.title("DOS for Uniform Stress vs Uniform Stress and Non-uniform Shear")
plt.xlabel("Energy (ev)")
plt.ylabel("Density of States")
plt.legend()
plt.savefig("dos_plot1v2.png")
plt.close()

# Plot uniform stress and Shear vs non-uniform stress non-uniform Shear 3 v 4
plt.plot(data["Energy (eV)"], data[f"Run #{49}"], label="Uniform Stress Plus Shear")
plt.plot(data["Energy (eV)"], data[f"Run #{73}"], label="Non-Uniform Stress Plus Shear")
plt.title("DOS for Non-Uniform Stress vs Uniform Stress and Shear On Principle Axes")
plt.xlabel("Energy (ev)")
plt.ylabel("Density of States")
plt.legend()
plt.savefig("dos_plot3v4.png")
plt.close()

# Plot uniform stress on principle axes vs stress and Shear 1 v 5
plt.plot(data["Energy (eV)"], data[f"Run #{1}"], label="Uniform Stress")
plt.plot(data["Energy (eV)"], data[f"Run #{97}"], label="Stress and Shear")
plt.title("DOS for Uniform Stress on Principle Axes vs Stress and Shear")
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
    label="Reference Axes on [0,0,0]",
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
    label="Reference Axes on [0,0,0]",
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
    dtype={"Run": float},
)

# copy the data out of the pandas dataframe to avoid modifying a data copy
dataEnergy = data["Run"].copy(deep=True)
# update run dataset to use correct percentage of max
for i in range(len(dataEnergy)):
    dataEnergy[i] = ((dataEnergy[i] - 1) * 4 / 9) / 4

plt.plot(dataEnergy, data["Energy (eV)"], label="Reference Axes on [0,0,0]")

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
    dtype={"Run": float},
)

# copy out of the dataframe to guarentee a slice
dataEnergy = data["Run"].copy(deep=True)
# update run dataset to use correct percentage of max
for i in range(len(dataEnergy)):
    dataEnergy[i] = ((dataEnergy[i] - 121) * 4 / 9) / 4

plt.plot(dataEnergy, data["Band Gap Energy (eV)"], label="Reference Axes on [0,0,0]")

plt.title("Band Gap for Increasing Uniform Stress (Max 4 GPA)")
plt.xlabel("Percentage of Max Stress")
plt.ylabel("Band Gap Energy (eV)")
plt.legend()
plt.savefig("bandGap_inc.png")
plt.close()
