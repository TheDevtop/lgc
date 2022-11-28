.TL
Little Green Clusters (Draft)
.AU
Thijs Haker
.AI
Fontys Hogescholen
.AB
Improving the computation/energy efficiency ratio,
by designing and implementing a service cluster,
and a collection of microservices.
With which small but rich network applications can be build.
.AE
.NH
Problem statement:
.PP
Since the dawn of computing in the 1960's, there is, and has been, a consistent tendency to run medium to high workloads on high power machines.
This is a reasonable way to do computing, except that there is significant power wastage when running medium sized jobs.
Moreover, a subset of these medium sized jobs (like compiling software), run for longer periods of time, increasing the overall power consumption.
.PP
Because of the risings in enegery cost, we would like to optimize our energy effeciency over computation.
And because compiling software is very widespread we will focus on this job.
.NH
Research Method:
.PP
This project is going to be researched via the engineering method.
.[
method
.]
Which is considered the applied-science counterpart to the scientific method.
This method is specified with prototype and research products in mind,
not with appliance or industrial products.
Conclusions and statements are made based on citations and sources.
Solutions are based on,
or in contrast with,
prior work.
.NH
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
.NH
Product research:
.PP
Since the breaktrough of cloud computing in the 2010's,
there have been several products that orchestrate and coordinate clusters.
In this section we will cover a few:
.NH 2
Kubernetes:
.PP
Kubernetes is a popular cluster orchestrator for application containers.
.[
k8sarch
.]
It's primarily made for datacenter size installations,
it has many features which makes it less suitable for our deployment situation.
However the overall architecture is highly structured and a good starting point for our system.
.NH 2 
Apache Mesos:
.PP
Mesos is a cluster scheduler,
it's possible to attach other cluster software (like Kubernetes) to it.
Altough it poses many interesting ideas about scheduling,
networking,
and data locality.
.[
mesosarch
.]
Many of these things are out of scope,
and primarily contribute to the increasing of complexity.
It could be an interesting idea to integrate LGC with Apache Mesos.
.NH 2
Plan 9:
.PP
This is the operating system developed after Bell Labs ended with UNIX.
It is not a cluster orchestrator,
but rather a distributed operating system.
.[
plan9arch
.]
What makes Plan 9 interesting is it's file protocol called 9P.
With 9P it is possible for CPU servers (compute nodes) to see a remote filesystem as if it where local.
Combining a remote filesystem with LGC would greatly simplify the operation of the cluster,
and solve the data locality problem.
Moreover,
in Plan 9,
the majority of operating system services are implemented as file servers that communicate via 9P.
.NH 2
Runners:
.PP
GitLab Runners are part of the GitLab CI/CD system,
with which build jobs can be executed.
It has many features which won't be considered for LGC,
but it's programming model/execution flow,
.[
runarch
.]
is simple enough to be implemented by a single person.
Altough it is not cluster orchistrator,
it can be used in a distributed manner.
.NH
Requirements:
.IP \1
The user can pick and choose which parts of the cluster to deploy.
.IP \2
The user can publish a service to the cluster.
.IP \3
The user can bind a published service to the API server.
.IP \4
The service can retrieve/store state from a centrally managed database.
.IP \5
The service can output it's messages to a centrally managed logging system.
.IP \6
The service can have any underlying system or platform.
.IP \6
The user can manage the cluster via the operating system console.
.NH
Evaluation:
.PP
There are two ways to solve this problem:
The first option is to build a cluster orchistrator that manages services,
and makes them exportable to the internet.
This will make the implementation of any build service secondary.
The second option is to build a multi-node build service,
this will be easier but will exclude the implementation of a full cluster orchestrator.
.PP
In the end I decided to design and implement a cluster orchistrator.
For both the reason that my interests are more aligned,
and the reason that there is more space to innovate.
.NH
Design & Implementation:
.PP
Little Green Cluster or
.B "LGC",
is a portable and elegant orchestration system for small computer clusters.
It lends its microservice design from Kubernetes,
.[
k8sarch
.]
and is written in Go both for portability,
and because Go was made for system and server software.
.[
golang
.]
The rest of this section can be found in a seperate document called:
.I "The design and implementation of LGC".
.NH
Conclusion:
.PP
Altough microservices and module oriented programming can improve efficiency by workload per unit.
The increase in communication and networking makes any such system more complex overall.
This isn't a problem when the communication is localized,
as with channels and pipes.
But in the era of multi-machine networking,
the communication becomes the problem.
.PP
This project started as way to efficiently compile and link source code.
But it has been generalized to provide a simple,
general purpose,
cluster orchistrator.
.PP
Therefore my conclusion and advice to any who reads this:
Efficient software is only as effective as the system it's part of.
And the hardware it's paired with,
whether it's storage, networking, or compute.
.NH 2
Future work:
.PP
Reimplement LGC based on the same principles as Plan 9
device servers.
.[
plan9dev
.]
This makes the system more monolithic,
but at the same time, it gives a simpler and more traditional interface.
