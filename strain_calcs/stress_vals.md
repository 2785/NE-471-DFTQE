# Current Script Information

Takes in stress from the a given reference axis, and returns the strain matrix (should change it give the sum of the sheer and axial stress in each direction). Axes are defined inline in the code. 

Input stress vals are defined in a separate file `stress.txt'. The syntax for the file is: 

data for 3x3 1
--
data for 3x3 2
--
...

Data for 3x3 is given in the following format
num,num,num,
num,num,num,
num,num,num,

## Reference Axis Currently scripted for

X=[1 0 0] Y=[0 1/sqrt(2) -1/sqrt(2)] Z=[0 1/sqrt(2) 1/sqrt(2)]

<!-- <img src="https://latex.codecogs.com/gif.latex?\left[1 0 0\right] t"/>  
<img src="https://latex.codecogs.com/gif.latex?Y=\left[0 \frac{1}{\sqrt{2}} -\frac{1}{\sqrt{2}}\right] t"/>  
<img src="https://latex.codecogs.com/gif.latex?Z=\left[0 \frac{1}{\sqrt{2}} \frac{1}{\sqrt{2}}\right] t"/>   -->