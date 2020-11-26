import math as m
import numpy as np
from pathlib import Path as p
from tarfile import TarFile as t
import os

def calculator( ref_Xaxis, ref_Yaxis, ref_Zaxis, stress, c_d , ai = [0.25,0.25,0.25]):
    '''
    Strain calculator given the stress factors for a FCC crystal, and the REF axis.

    Inputs: 
        ref_Xaxis - reference axis for the X axis (list length 3)
        ref_Yaxis - reference axis for the Y axis (list length 3)
        ref_Zaxis - reference axis for the Z axis (list length 3)
        stress - stress applied in the given reference axis (3x3 nested list)
        c_d - stress constants for the solid to be solved for (since FCC this is 3membered list)
        ai - this is the latice positions of the corner atoms (1/4,1/4,1/4).
    Output:
        Outputs a tuple of the new lattice positions of the unit cell. 
        (atom1, atom2).
        atom1 - List of length 3 indicating the lattice position of atom 1. 
        atom2 - List of length 3 indicating the lattice positon of atom 2. 
        Assuming formula is: Atom1Atom2
    '''
    # Setting up state
    row = 3
    col = 3
    size = (row,col)
    s_col = 6
    strain = np.zeros(size)
    stress_col = np.zeros((s_col,1))

    a = np.transpose( np.array([ref_Xaxis,ref_Yaxis,ref_Zaxis]) )
    ## Figure out the stress 3x3 system. 
    r1 = [c_d[0],c_d[1],c_d[1],0,0,0]
    r2 = [c_d[1],c_d[0],c_d[1],0,0,0]
    r3 = [c_d[1],c_d[1],c_d[0],0,0,0]
    r4 = [0,0,0,c_d[2],0,0]
    r5 = [0,0,0,0,c_d[2],0]
    r6 = [0,0,0,0,0,c_d[2]]
    c_unit_d = np.array([r1,r2,r3,r4,r5,r6])

    c_unit = c_unit_d * 0.1 #conversion to Pa

    stress_unit = np.zeros(size)
    #assuming square.
    for i in range(row):
        for j in range(col):
            for k in range(col):
                for l in range(row):
                    stress_unit[i][j] += a[i][k] * a[j][l] * stress[k][l]


    # 3x3 -> col 
    for i in range(s_col):
        if i <= 2:
            stress_col[i] = stress_unit[i][i]
        elif 3 <= i <= 4:
            stress_col[i] = 2*stress_unit[i%2][2]
        else:
            stress_col[i] = 2*stress_unit[0][1]

    #tablulate the elastic constant matrix using FCC lattice

    c_inv = np.linalg.inv(c_unit)
    strain_col = np.matmul(c_inv, stress_col)

    #col -> 3x3 
    for i in range(s_col):
        if i <= 2:
            strain[i][i] = strain_col[i]
        elif 3 <= i <= 4:
            strain[i%2][2] = strain_col[i]/2
            strain[2][i%2] = strain_col[i]/2
        else:
            strain[0][1] = strain_col[i]/2
            strain[1][0] = strain_col[i]/2
    strain_np = np.array(strain)
    atom2 = [ai[0],ai[1],ai[2]]
    atom1 = [0,0,0]
    for r_i in range(row):
        for c_i in range(col):
            if r_i == c_i :
                atom1[r_i] += strain_np[r_i,c_i]
            else:
                atom2[c_i] += strain_np[r_i,c_i]
    # return (ai[0] + sum(strain_np[:,0]), ai[1] + sum(strain_np[:,1]), ai[2] + sum(strain_np[:,2]) ) #this only changes the corner atoms
    return (atom1,atom2) #this changes the entire structure
    # atom2 represents those on the corner
    # atom1 represents those on the face and center
    # Output strain given in the 1 0 0 ref axis. 

# units of the below are dyn/cm^2
c11_d = 11.90e11
c12_d = 5.34e11
c44_d = 5.96e11

c = [c11_d, c12_d, c44_d]

axis_vals = [[0, 1/m.sqrt(2), 1], [0, -1/m.sqrt(2), 1], [1/m.sqrt(6), 2/m.sqrt(6), 3/m.sqrt(6)], [-3/m.sqrt(8), -1/m.sqrt(8), -4/m.sqrt(8)]]
folder = p.cwd() / "n471-proj-carrot" / "strain_calcs"
# Modify folder to match your dir, currently on my dir. 
name_cnt = 1
compound = "Gallium Arsenide diamond structure\n"
zipfile = folder / "test_files.tar.gz" #creates tarball file
# zipper = gz.open(zipfile, 'w')
zipper = t.open(zipfile, 'w:gz')
for g, x_val in enumerate(axis_vals):
    ref_X = x_val
    for h, y_val in enumerate(axis_vals):
        if (h != g) and (x_val != y_val): 
            ref_Y = y_val
        else:
            continue
        for j, z_val in enumerate(axis_vals):
            # if j != g and j != h: #or z_val == y_val: to take out duplicates
            if (j != g) and (z_val != y_val): 
                ref_Z= z_val
                #set up some system to define the stress values. 

                # file w/ stress matrices
                # while file.readline() -> read the first line
                # do all the stuff
                # have mutiple zip folders

                stresser = folder / "stress.txt"
                stress = []
                with open(stresser, "r") as cs:
                    for o, line in enumerate(cs.readlines()):
                        row_val = line
                        temp = []
                        if "-" not in line:
                            while row_val != "\n":
                                delim = row_val.index(',')
                                temp += [float(row_val[:delim])*1e9,]
                                row_val = row_val[delim+1:]
                            stress += [temp,]
                        else:
                            # do the calculations
                            (ga, ars) = calculator(ref_X, ref_Y, ref_Z, stress, c)
                            # store the values in text file in the format for nanohub
                            # open file
                            fname = "nano_hub_test_"+str(name_cnt)+".txt"
                            filename = fname
                            f = open(filename, "w")
                            f.write("2\n"+compound)
                            f.write("Ga " + str(ga[0]) + " " + str(ga[1]) + " " + str(ga[2]) + "\n")
                            f.write("As " + str(ars[0]) + " " + str(ars[1]) + " " + str(ars[2]) + "\n")
                            f.write("---\n")
                            f.write("X refaxis: " + str(x_val) + "\nY ref axis " + str(y_val) + "\nZ refaxis " + str(z_val) + "\n")
                            f.close()
                            cmd_clean = "rm " + filename.__str__()
                            zipper.add(filename) # Adds file to tarball
                            x_rm = os.system(cmd_clean) # Deletes original file
                            name_cnt += 1
                            stress = []
                    cs.close()
            else:
                continue
latticefile = "lattice.txt"
with open(latticefile, 'w') as l:
    l.write("5.6325")
l.close()
zipper.add(latticefile)
cmd = "rm " + latticefile.__str__()
x_rm = os.system(cmd)
zipper.close()
