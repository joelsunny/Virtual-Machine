# function for finding maximum of two numbers
proc max:
load 5
load 4
isub
iconst 0
ilt
brf 4
load 4 
iret
load 5
iret
end

iconst 13
iconst 19
call max 2 # max label will be replaced by assembler
print
halt