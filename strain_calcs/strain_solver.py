import math as m
import numpy as np


# units of the below are dyn/cm^2
c11_d = 11.90e11
c12_d = 5.34e11
c44_d = 5.96e11

# Setting up state
row = 3
col = 3
size = (row,col)
s_col = 6
stress = np.zeros(size)
strain = np.zeros(size)
stress_col = np.zeros((s_col,1))
ref_Xaxis = [1,0,0]
ref_Yaxis = [0,1/m.sqrt(2),-1/m.sqrt(2)]
ref_Zaxis = [0,1/m.sqrt(2),1/m.sqrt(2)]
# a = np.zeros(size)
# v = ['X','Y','Z']
# for i in range(row):
#     row_val = input('Enter the reference axis of ' + v[i] + ' information separated by commas\n')
#     temp = []
#     while len(row_val) > 1:
#         delim = row_val.index(',')
#         temp += [float(row_val[:delim]),]
#         row_val = row_val[delim+1:]
#     temp += [float(row_val),]
#     a[i] = temp

a = np.transpose( np.array([ref_Xaxis,ref_Yaxis,ref_Zaxis]) )
## Figure out the stress 3x3 system. 

c_unit_d = np.array([[c11_d,c12_d,c12_d,0,0,0],[c12_d,c11_d,c12_d,0,0,0],[c12_d,c12_d,c11_d,0,0,0],[0,0,0,c44_d,0,0],[0,0,0,0,c44_d,0],[0,0,0,0,0,c44_d]])

c_unit = c_unit_d * 0.1 #concersion to Pa

for i in range(row):
    row_val = input('Enter the row' + str(i+1) + ' elements of stress in the reference axis specified separated by commas\n')
    temp = []
    while len(row_val) > 1:
        delim = row_val.index(',')
        temp += [float(row_val[:delim]),]
        row_val = row_val[delim+1:]
    temp += [float(row_val),]
    stress[i] = temp

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

# Output strain given in the 1 0 0 ref axis. 