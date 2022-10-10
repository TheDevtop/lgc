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
A much simpler option is to have a set of centrally orchestrated computers,
here the only bottleneck is the supported architectures of the orchestration software.
This type of software is known as "cluster orchestation" software.
.SH
Product research:
.PP
Since the breaktrough of cloud computing in the 2010's,
there have been several products that orchestrate and coordinate clusters.
In this section we will cover a few.
.PP
.B "Kubernetes:"
Kubernetes is a popular cluster orchestrator for application containers.
.[
k8sarch
.]
It's primarily made for datacenter size installations,
it has many features which makes it less suitable for our deployment situation.
However the overall architecture is highly structured and a good starting point for our system.
.PP
.B "Apache Mesos:"
Mesos is a cluster scheduler,
it's possible to attach other cluster software (like Kubernetes) to it.
.[
mesosarch
.]
Altough it poses many interesting ideas about scheduling,
networking,
and data locality.
But many of these things are out of scope.
It would be an interesting idea to integrate LGC with Apache Mesos,
this would be considered future work tough.
.PP
.B "Plan 9:"
This is the operating system developed after Bell Labs ended with UNIX.
It is not a cluster orchestrator,
but rather a distributed operating system.
.[
plan9arch
.]
What makes Plan 9 interesting is it's file protocol called 9P.
With 9P it is possible for CPU servers (compute nodes) to see a remote filesystem as if it where a local filesystem.
And I would like this idea of a central filesystem to be part of LGC.
This also solves the data locality problem.
.SH
Requirements:
.PP
In this section we will list the requirements,
and explain them in further detail.
