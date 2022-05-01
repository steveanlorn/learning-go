# Concurrency Pattern

## #1 Wait for Result

The main idea behind Wait for Result pattern is to have:
- a channel that provides a signaling semantics
- a goroutine that does some task
- a goroutine that waits for that task to be done

**Scenario and Example**   
Code: `01_wait_for_result`  

You are a manager and have a new employee.
The new employee understands what to do without being told.
You are waiting for the task to be done.
The amount of time you wait is unknown 
because you need a guarantee that the done task is notified to you.

**Note**
- The order of `fmt.Println()` executions are not guaranteed.

---

## #2 Fan Out

The main idea behind Fan Out Pattern is to have:

- a channel that provides a signaling semantics
  - channel can be buffered, so we don't wait on immediate receive confirmation
- a goroutine that starts multiple (other) goroutines to do some task
- a multiple goroutines that do some task and use signaling channel to signal that the task is done


**Scenario and Example**   
Code: `02_fan_out`

You are a manager and have multiple employees.
The new employees understand what to do without being told.
You are waiting for the task to be done.
The amount of time you wait is unknown
because you need a guarantee that all the done tasks are notified.

**Note**
- Fan out pattern can give a lot of load on the system & external system at the same time.

---

## #3 Wait for Task

The main idea behind Wait for Task pattern is to have:
- a channel that provides a signaling semantics
- a goroutine that **waits for task**, so it can do some task
- a goroutine that sends task to the previous goroutine

**Scenario and Example**   
Code: `03_wait_for_task`

You are a manager and have a new employee.
The new employee does not know what to do and waits for you to tell.
You prepare the task and send it to the new employee.
The amount of time the new employee wait is unknown
because you need a guarantee that the task you are sending is received by the employee.

**Note**
- Wait for task pattern is the base pattern for the pool pattern.


---

## #4 Pooling

The main idea behind Pooling pattern is to have:
- a channel that provides a signaling semantics
  - unbuffered channel is used to have a guarantee a goroutine has received a signal
- multiple goroutines that pool that channel for task
- a goroutine that sends task via channel


**Scenario and Example**   
Code: `04_pooling`

You are a manager and have multiple employees.
The new employees do not know what to do and wait for you to tell them.
You prepare the task and send it to them.
The first available employee takes the job.
Once the employee finishes the job, he is waiting for the new job.
The amount of time you wait for any given employee to take your task is unknown because you need a
guarantee that the task you send is received by the employee.


**Note**
- Use unbuffered channel to guarantee the worker has received the task and working on it.

---

## #5 Fan Out Semaphore
The main idea behind Fan Out Semaphore Pattern is to have:

Everything we had in the Fan Out Pattern:
- a buffered channel that provides a signaling semantics
- a goroutine that starts multiple (child) goroutines to do some task
- a multiple (child) goroutines that do some task and use signaling channel to signal the task is done

PLUS the addition of a:
- new semaphore channel used to restrict the number of child goroutines that can be schedule to run


**Scenario and Example**   
Code: `05_fan_out_semaphore`

You are a manager and have multiple employees.
The new employees understand what to do without being told.
You are waiting for the task to be done.
The amount of time you wait is unknown
because you need a guarantee that all the done tasks are notified.
**BUT** computer for the employee to do the task only available 10.
Other employees need to wait until 10 employees finish the task and free the computer.

**Note**
- Good use case for this pattern would be batch processing, where we have some amount of task to do, but we want to limit the number of active executors at any given moment.

---

## #6 Fan Out Bounded

The main idea behind Fan Out Bounded Pattern is to have a limited number of goroutines that will do the task. We have:
- a fixed number of worker goroutines
- a manager goroutine that creates/reads the task and sends it to the worker goroutines
- a buffered channel that provides signaling semantics
  - used to notify worker goroutines about available task

**Scenario and Example**   
Code: `06_fan_out_bounded`

You are a manager and have limited employees.
The new employees do not know what to do and wait for you to tell them.
You prepare the task and send it to them.
Once all the employee are full, you put the next task on queue.
If the queue is full, you can not put more task and wait until the queue is reduced.
The amount of time you wait for any given employee to take your task is unknown because you need a
guarantee that the task you send is received by the employee.

**Note**
- Good use case for batch process where we have limited number of executors.

---

## #7 Drop

The main idea behind Drop Pattern is to have a limit on the amount of task that can be done at any given moment.
We have:
- a buffered channel that provides signaling semantic
- a number of worker goroutines
- a manager goroutine that:
  - takes the task and sends it to the worker goroutine
  - if there is more task than worker goroutines can process and buffered channel is full, manager goroutine will drop the task

**Scenario and Example**   
Code: `07_drop`

You are a manager and have employees. 
You prepare the task and send it to the employee.
You don't wait for the employee to take the task.
If the employee is not ready then you drop the task and try again with the next task.


**Note**
- Good use case for reducing pressure on the system. Only allow request that within the capacity.

---

## #8 Cancellation

The main idea behind the Cancellation Pattern is to have a limited amount of time to perform task. 
If limit is reached, the task is ignored.

We have:
- a context with specified timeout
- a buffered channel that provides signaling semantic
- a worker goroutine that does the task
- a manager goroutine that waits on (which comes first):
  - worker goroutine signal (that the task is completed)
  - context timeout signal

**Scenario and Example**   
Code: `08_cancellation`

You are a manager and have employees.
The employees are working with their task.
You are not waiting to wait forever for the employee to finish their task.
They have specified amount of time.

**Note**
- Good use case for this pattern is any request to a remote service, e.g. database request, API request or whatever request that can block. Since we don't want our request to block forever, we use timeout to cancel it.