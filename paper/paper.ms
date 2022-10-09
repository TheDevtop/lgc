.TL
Little Green Clusters (Draft)
.AU
Thijs Haker
.AI
Fontys Hogescholen
.SH
Problem statement:
.PP
Since the dawn of computing in the 1960's, there is, and has been, a consistent tendency to run medium to high workloads on high power machines.
This is a reasonable way to do computing, except that there is significant power wastage when running medium sized jobs.
Moreover, a subset of these medium sized jobs (like compiling software), run for longer periods of time, increasing the overall power consumption.
.PP
Because of the risings in enegery cost, we would like to optimize our energy effeciency over computation.
And because compiling software is very widespread we will focus on this job.
.SH
Background research:
.PP
Around 2008 the industry released their CPU's with the highest clock rate, 
and going any higher is considered not worth the cost in terms of power consumed and heat dissipated.
.[
cpufreq
.]
There are methods for the mitigation of power consumption.
But with the advent of multiple cores,
these methods yielded diminished returns.
.[
dvfs
.]
However,
the combination of multiple low powered CPU cores stike an optimal balance between effeciency and computation.
Contemporary computers with this design include the Raspberry Pi
.[
pi4
.]
and the Pine64.
.PP
Because our goal is to compile software,
we are required to assume the existence of multiple OS/Architecture combinations.
.[
gccarch
.]
We could solve this problem via emulation,
however this can (depending on the workload) introduce significant overhead,
.[
overhead
.]
even more so on a low powered machine.

