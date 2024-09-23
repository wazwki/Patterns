# Golang patterns implementation

## **Creational Patterns**
These patterns help to efficiently create objects by managing the complexity of object creation.

---

### 1. **Singleton** - Creational Pattern
**Purpose:** Ensures that a class has only one instance and provides global access to it.

**Usage:**
- To control access to a single instance (e.g., logging, database connections, configurations).

---

### 2. **Factory Method** - Creational Pattern
**Purpose:** Defines an interface for creating an object but allows subclasses to alter the type of objects that will be created.

**Usage:**
- When you need to encapsulate the process of object creation without depending on specific classes.

---

### 3. **Abstract Factory** - Creational Pattern
**Purpose:** Provides an interface for creating families of related or dependent objects without specifying their concrete classes.

**Usage:**
- To create objects that need to work together (e.g., UI components for different platforms).

---

### 4. **Builder** - Creational Pattern
**Purpose:** Allows step-by-step creation of complex objects.

**Usage:**
- For constructing objects with different configurations (e.g., building complex products or objects with many parameters).

---

### 5. **Prototype** - Creational Pattern
**Purpose:** Creates objects by copying an existing object (prototype).

**Usage:**
- When creating an object is expensive or complicated, and it's easier to clone an existing one.

---

## **Structural Patterns**
These patterns help organize relationships between objects to achieve flexibility and extensibility.

---

### 6. **Adapter** - Structural Pattern
**Purpose:** Allows objects with incompatible interfaces to work together.

**Usage:**
- When you need to use an object with an incompatible interface.

---

### 7. **Decorator** - Structural Pattern
**Purpose:** Dynamically adds new responsibilities to an object without changing its structure.

**Usage:**
- To add functionality to an object on the fly (e.g., adding visual effects to UI components).

---

### 8. **Facade** - Structural Pattern
**Purpose:** Provides a simple interface to a complex system.

**Usage:**
- When you need to simplify interaction with multiple subsystems or libraries.

---

### 9. **Composite** - Structural Pattern
**Purpose:** Allows grouping objects into tree structures and working with them as if they were individual objects.

**Usage:**
- To represent hierarchies of objects (e.g., file system structure).

---

### 10. **Proxy** - Structural Pattern
**Purpose:** Provides a surrogate or placeholder for another object to control access to it.

**Usage:**
- To control access, caching, or lazy initialization of objects.

---

### 11. **Bridge** - Structural Pattern
**Purpose:** Separates an abstraction from its implementation, allowing them to evolve independently.

**Usage:**
- When you need to separate implementation and abstraction for independent modification.

---

### 12. **Flyweight** - Structural Pattern
**Purpose:** Reduces the number of objects created by reusing existing ones.

**Usage:**
- When you need to manage many small objects (e.g., characters in a text editor).

---

## **Behavioral Patterns**
These patterns help organize effective communication between objects, ensuring flexibility and ease of collaboration.

---

### 13. **Strategy** - Behavioral Pattern
**Purpose:** Encapsulates a family of algorithms, making them interchangeable.

**Usage:**
- When you need to choose between different behavior options depending on conditions.

---

### 14. **Observer** - Behavioral Pattern
**Purpose:** Defines a one-to-many dependency, where the state change of one object notifies all dependent objects.

**Usage:**
- To implement event subscription mechanisms (e.g., newsletters).

---

### 15. **Command** - Behavioral Pattern
**Purpose:** Encapsulates a request as an object, allowing parameterization of clients with different requests.

**Usage:**
- To handle operations as objects (e.g., undo/redo systems).

---

### 16. **State** - Behavioral Pattern
**Purpose:** Allows an object to change its behavior when its internal state changes.

**Usage:**
- When an object needs to alter its behavior based on internal state (e.g., ATM or vending machine).

---

### 17. **Template Method** - Behavioral Pattern
**Purpose:** Defines the skeleton of an algorithm in the base class and allows subclasses to override specific steps.

**Usage:**
- To create algorithms where certain steps may vary in subclasses.

---

### 18. **Visitor** - Behavioral Pattern
**Purpose:** Allows adding new operations to objects without changing their classes.

**Usage:**
- When you need to perform various operations on objects of different classes without altering their code.

---

### 19. **Mediator** - Behavioral Pattern
**Purpose:** Simplifies communication between objects by ensuring centralized interaction through a mediator.

**Usage:**
- To reduce dependencies between objects and improve interaction through a single object.

---

### 20. **Chain of Responsibility** - Behavioral Pattern
**Purpose:** Allows passing a request along a chain of handlers until one of them processes it.

**Usage:**
- When a request needs to be processed by several objects.

---

### 21. **Interpreter** - Behavioral Pattern
**Purpose:** Defines a grammar for a language and interprets expressions in that language.

**Usage:**
- When you need to interpret and analyze a specific language or code.

---

### 22. **Memento** - Behavioral Pattern
**Purpose:** Allows saving and restoring the state of an object without violating its encapsulation.

**Usage:**
- When it’s necessary to save the state of an object for later restoration (e.g., undo/redo systems).

---

## **Concurrency Patterns**
These patterns help organize interaction between goroutines and ensure efficient execution of parallel tasks.

---

### 23. **Worker Pool**
**Purpose:** Organizes a pool of goroutines (workers) to process tasks in parallel.

**Usage:**
- When you need to distribute many independent tasks among a limited number of workers to manage resources.

---

### 24. **Fan-In / Fan-Out**
**Purpose:** Executes multiple tasks in parallel (Fan-Out) and collects results from several goroutines into one (Fan-In).

**Usage:**
- When you need to perform several operations in parallel and collect their results in one place.

---

### 25. **Pipeline**
**Purpose:** Breaks a task into several stages, where each stage is executed in a separate goroutine and passes the result to the next stage via channels.

**Usage:**
- When a task can be divided into several sequential steps, each of which can be executed independently.

---

### 26. **Select Statement**
**Purpose:** Allows managing multiple channels, waiting for data from any of them.

**Usage:**
- When you need to react to events or data coming from different sources (channels).

---

### 27. **Cancellation**
**Purpose:** Allows stopping the execution of goroutines that are no longer needed using a cancellation mechanism (`context.Context`).

**Usage:**
- When you need to terminate several goroutines in case of an error or when the task is no longer required.

---

### 28. **Timeouts**
**Purpose:** Allows setting a maximum wait time for a goroutine or an external resource response.

**Usage:**
- When it’s important to limit the time for task execution and terminate it in case of exceeding the allowed time.

---

### 29. **Semaphore**
**Purpose:** Limits the number of goroutines running simultaneously using a counter (`semaphore`) to manage access to shared resources.

**Usage:**
- When you need to control the number of concurrent operations, such as limiting load on resources.

---
