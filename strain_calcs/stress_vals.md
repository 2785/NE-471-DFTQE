# Current Script Information

Takes in stress from the a given reference axis, and returns the strain matrix (should change it give the sum of the sheer and axial stress in each direction). Axes are defined inline in the code. 

Input stress vals are defined in a separate file `stress.txt'. The syntax for the file is: 

data for 3x3 1
{----}
data for 3x3 2
{----}
etc

Data for 3x3 is given in the following format
num,num,num,
num,num,num,
num,num,num,


{----} lines just need to contain the '-' character to be skipped. 