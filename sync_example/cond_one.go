package syncexample


func CondOne() {

}


// The sync.Cond type provides a way to create and manage condition 
// variables. It has three main methods:

// Wait(): This method causes the calling goroutine to wait until another 
// goroutine signals the condition variable. When the goroutine calls Wait(), 
// it releases the associated lock and suspends execution until another goroutine 
// calls Signal() or Broadcast() on the same sync.Cond variable.

// Signal(): This method wakes up one goroutine waiting on the condition 
// variable. If multiple goroutines are waiting, only one of them is awakened. 
// The choice of which goroutine gets awakened is arbitrary and not guaranteed.

// Broadcast(): This method wakes up all goroutines waiting for the condition
// variable. When Broadcast() is called, all waiting goroutines are awakened 
// and can proceed.