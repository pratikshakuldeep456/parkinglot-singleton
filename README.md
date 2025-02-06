
# Designing a Parking Lot System

## Requirements
1. The parking lot should have multiple levels, each level with a certain number of parking spots.
2. The parking lot should support different types of vehicles, such as cars, motorcycles, and trucks.
3. Each parking spot should be able to accommodate a specific type of vehicle.
4. The system should assign a parking spot to a vehicle upon entry and release it when the vehicle exits.
5. The system should track the availability of parking spots and provide real-time information to customers.
6. The system should handle multiple entry and exit points and support concurrent access.



# Design Patterns can be used:
1. Singleton Pattern: Ensure only one instance of the parking lot system exists.
2. Strategy Pattern: Flexible parking spot allocation based on vehicle type.
3. Observer Pattern: Real-time updates to notify customers of spot availability.
4. Factory Pattern: Create different vehicle types dynamically.
5. Mutex/Concurrency Control: Handle concurrent access for multiple entries/exits.
