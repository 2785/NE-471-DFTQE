import math as m
import numpy as np
from pathlib import Path as p
from zipfile import ZipFile as z

def calculator( ref_Xaxis, ref_Yaxis, ref_Zaxis, stress, c_d , ai = [0.25,0.25,0.25]):
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
    for i in range(size[0]):
        for j in range(size[1]):
            for k in range(size[0]):
                for l in range(size[1]):
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
    return (ai[0] + sum(strain_np[:,0]), ai[1] + sum(strain_np[:,1]), ai[2] + sum(strain_np[:,2]) )
    # Output strain given in the 1 0 0 ref axis. 

# units of the below are dyn/cm^2
c11_d = 11.90e11
c12_d = 5.34e11
c44_d = 5.96e11

c = [c11_d, c12_d, c44_d]

# ref_Xaxis = [1,0,0]
# ref_Yaxis = [0,1/m.sqrt(2),-1/m.sqrt(2)]
# ref_Zaxis = [0,1/m.sqrt(2),1/m.sqrt(2)]

axis_vals = [[0, 0, 1], [0, 1/m.sqrt(2), 1], [0, -1/m.sqrt(2), 1], [1/m.sqrt(6), 2/m.sqrt(6), 3/m.sqrt(6)], [-3/m.sqrt(8), -1/m.sqrt(8), -4/m.sqrt(8)]]
folder = p.cwd() / "n471-proj-carrot" / "strain_calcs"
# Modify folder to match your dir, currently on my dir. 
name_cnt = 1
compound = "Gallium Arsenide diamond structure\n"
zipfile = folder / "files.zip"
# zipper = z.open(zipfile.__str__(), 'a')
for g, x_val in enumerate(axis_vals):
    ref_X = x_val
    for h, y_val in enumerate(axis_vals):
        if h != g: 
            ref_Y = y_val
        else:
            continue
        for j, z_val in enumerate(axis_vals):
            if j != g and j != h: #or z_val == y_val: to take out duplicates
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
                                temp += [float(row_val[:delim]),]
                                row_val = row_val[delim+1:]
                            stress += [temp,]
                        else:
                            # do the calculations
                            sol = calculator(ref_X, ref_Y, ref_Z, stress, c)
                            # store the values in text file in the format for nanohub
                            # open file
                            fname = "nano_hub_test_"+str(name_cnt)+".txt"
                            filename = folder / fname
                            f = open(filename, "w")
                            f.write("2\n"+compound)
                            f.write("Ga 0.0 0.0 0.0\n")
                            f.write("As " + str(sol[0]) + " " + str(sol[1]) + " " + str(sol[2]) + "\n")
                            f.write("---\n")
                            f.write("X refaxis: " + str(x_val) + "\tY ref axis " + str(y_val) + "\tZ refaxis " + str(z_val) + "\n")
                            f.close()
                            # zipper.write(filename)
                            # p.unlink(filename)
                            name_cnt += 1
                            stress = []
                            # make sure that the number of stress matrices equals the number of axis vals. 
                    cs.close()
            else:
                continue
# zipper.close()
#zip *.txt nano_hub_files.zip
