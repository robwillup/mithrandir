# Introduction

## Exercises

### 1.1-1:
Give a real-world example that requires sorting or a real-world example that requires computing a convex hull.

    A real-world example for sorting is a e-commerce web page that allows the user to sort products by name
    A real-world example the requires calculating convex hull is to find out how much wire is needed for a fence.

### 1.1-2
Other than speed, what other measures of efficiency might one use in a real-world setting?

    Another measure of efficiency is the usage of resources. A fast car that burns too much fuel is not efficient.

### 1.1-3
Select a data st ructure that you have seen previously, and discuss its strengths and limitations.

    Regarding dictionary or map, its strengths are quick retrieval, quick insertion. Limitations are slow look up.

### 1.1-4
How are the `shor test-path` and the `traveling-salesman` problems similar? How are they different?

    Both try to find the smallest distance. But shortest-path just tries to get from one point to another, while
    traveling-salesman has to visit all the points.

### 1.1-5
Come up with a rea l-world problem in which only the best solution will do. Then come up with one in which a solution
that is "approximately" the best is good enough.

    One that requires the best solution only is an app that calcuates heart bits of patients in real time. One that
    is ok with approximately is an app that's responsible for sorting records in the background (although it could be
    argued that it would be wasting money if running on the cloud or energy and resources if onprem)

#### Answers from GPT-5-Codex

1.1-1

    Sorting example: An online retailer lets shoppers sort products by ascending price so they can quickly find the cheapest option.
    Convex hull example: Urban planners compute the convex hull of GPS points marking existing buildings to determine the minimal perimeter needed for a protective barrier.

1.1-2

    Efficiency can also be measured by memory consumption, energy usage, financial cost, development effort, and environmental impact.

1.1-3

    Hash map (dictionary)

    Strengths: Average-case O(1) insertion, lookup, and deletion; flexible key types; ideal for fast key-based access.
    Limitations: No inherent ordering of keys; worst-case operations degrade to O(n) with poor hashing; higher memory overhead due to buckets.

1.1-4

    Both problems search for paths with minimal total weight. The shortest-path problem finds the minimum-cost route between specific nodes, whereas the traveling-salesman problem seeks the minimum-cost tour that visits every node exactly once and returns to the start, making it combinatorially harder.

1.1-5

    Exact solution required: A collision-avoidance system for autonomous vehicles must compute the optimal evasive maneuver to guarantee safety.
    Approximate solution acceptable: A ridesharing app can use near-optimal driver-passenger matching to balance wait times and fuel use, accepting slight inefficiencies for speed and scalability.

## Algorithms as a technology

Computing time is therefore a bounded resource, and so is space in memory. You should use these resources wisely, and
algorithms that are efficient in terms of time or space will help you do so.

We should consider algorithms, like computer hardware, as a technology. Total system performance depends on choosing
efficient algorithms as much as on choosing fast hardware.

Although some applications do not explicitly require algorithmic content at the application level (such as some simple,
Web-based applications), many do. Even an application that does not require algorithmic content at the application
level relies heavily upon algorithms. Does the application rely on fast hardware? The hardware design used algorithms.
Does the application rely on graphical user interfaces? The design of any GUI relies on algorithms. Does the
application rely on networking? Routing in networks relies heavily on algorithms. Was the application written in a
language other than machine code? Then it was processed by a compiler, interpreter, or assembler, all of which make
extensive use of algorithms. Algorithms are at the core of most technologies used in contemporary computers.

Having a solid base of algorithmic knowledge and technique is one characteristic that separates the truly skilled
programmers from the novices.

### Exercises

#### 1.2-1
Give an example of an application that requires algorithmic content at the application level, and discuss the function
of the algorithms involved.

My answer:

    An application to enhance image resolution. The function of the algorithm is to optimize the quality of the image
    in the most effecient way. The algorithms used include:
    * Interpolation: bicubic or bilinear interpolation estimates pixel values between existing ones to upscale images.
    * Convolutional Neural Networks: Deep learning models like SRCNN, ESRGAN learn patterns from high-resolution
      training data to intelligently fill in missing details.
    * Edge detection and sharpening: identifies edges and enhances them to reduce blur after upscaling.
    * Optimization algorithms: Gradient descent trains neural networks to minimize reconstruction error between upscaled
      and original high-res images.

    Modern approaches rely heavily on machine learning to produce higher-quality results than traditional mathematical
    interpolation.



