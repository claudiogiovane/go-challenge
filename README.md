# Solution report Go challenge


# Task
Write an HTTP service that exposes an endpoint "/numbers". This endpoint receives a list of URLs though GET query parameters. The parameter in use is called "u". It can appear more than once.

http://yourserver:8080/numbers?u=http://example.com/primes&u=http://foobar.com/fibo
When the /numbers is called, your service shall retrieve each of these URLs if they turn out to be syntactically valid URLs. Each URL will return a JSON data structure that looks like this:

## { "numbers": [ 1, 2, 3, 5, 8, 13 ] } 

The JSON data structure will contain an object with a key named "numbers", and a value that is a list of integers. After retrieving each of these URLs, the service shall merge the integers coming from all URLs, sort them in ascending order, and make sure that each integer only appears once in the result. The endpoint shall then return a JSON data structure like in the example above with the result as the list of integers.

The endpoint needs to return the result as quickly as possible, but always within 500 milliseconds. It needs to be able to deal with error conditions when retrieving the URLs. If a URL takes too long to respond, it must be ignored. It is valid to return an empty list as result only if all URLs returned errors or took too long to respond.

<br>

# Completion Conditions

Solve the task described above using Go. Only use what's provided in the Go standard library. The resulting program must run stand-alone with no other dependencies than the Go compiler.

Document your source code, both using comments and in a separate text file that describes the intentions and rationale behind your solution. Also write down any ambiguities that you see in the task description, and describe you how you interpreted them and why. If applicable, write automated tests for your code.

For testing purposes, you will be provided with an example server that, when run, listens on port 8090 and provides the endpoints /primes, /fibo, /odd and /rand.

<br>

# Interpretation of the problem

Initially I couldn't quite understand how to test the solution. I found that urls should be those described in the problem, like http://example.com/primes and so on. After a few minutes I realized that http: // localhost: 8090/primes, for example, would return the values as quoted in the task. So I realized that in the same localhost, but on port 8080, I should upload my service.

**Development assumptions**

Before even starting to code, i had already decided that i would make a simplistic approach as possible, using the resources provided by the language itself and build the code clean without losing, however, the focus of the performance, since the task required a maximum response time of 500ms.

**Project Structure** 

I started with the main and handler functions and decided that the main file would contain only these two. The data processing functions would be in a separate file, and the more general functions would also be in a third file. Thus, the code would be more readable and easy to understand.

*handler*

The handler was the most difficult to implement since it was the core of the service. First, I built it as simply as possible, just handling the parameters (URLs) and retrieving the data.
In this process, I stated the functions that would be required for sorting and merging the numbers, I left the handler aside and begin to implement these functions in a separate file as I had originally planned. I'll get back to the handler in the end.

*Reorg function*

The first function to be implemented was the Reorg, which is called two other functions, the sort function and the one that eliminates repetitions, both very simple. In the sort function, there is only one error handling and the call of the sort.Ints () function, following, thus the initial premise not to rummage the code too much and to use to the maximum what the language itself has to offer. For the function of removing the repetitions, a simple loop is passing through the slice and removing the repetitions.

*Mergenumbers function*

The merge function is also extremely simple. In it, we join the slices passed as parameter and then we call again the functions of sorting and elimination of repetition.

*Generic functions and validations*

After implementing the data manipulation functions, I begin to implement the generic functions of the service that were the validation of the URLs, JSON Parser and finally the one that retrieved the numbers of the JSON received. All of these were implemented in a separate file.

*Ending the handler*

After all the accessory functions were done, I focused on the handler again. Until then, it had not implemented any solution that gave performance to the service, which resulted in exceeding the maximum time required almost 100% of the time. So, It was necessary that the processing of the URLs be done in parallel. That's when I add the goroutine and channel, which make concurrency processing extremely easy and effective. Thus, I opened a channel for the traffic of the sets of numbers when retrieving then and closing it at the end of the merge. After that, the performance of the service reached the satisfactory level and I was finally able to implement the automated tests.
