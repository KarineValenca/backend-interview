
### What metrics could be interesting to observe to ensure app integrity and stability?
  To ensure the app's integrity and stability, I consider metrics like error rate, latency, and traffic really useful. With the error rate metric, it's possible to know how many errors the users are receiving. The latency metric allows to know how much time an endpoint is taking to load. And the traffic metric allows to know the number of accesses in an endpoint. With these three metrics, it's possible to have an overview of the application's health. 

### What kind of alerts based on those metrics could we use here? What critical conditions should we look at?
  To create alerts based on metrics, it's important to know the application's behavior. It's not possible to say that an error rate above 50% in an endpoint is problematic, if we don't have knowledge of the application behavior. My approach to create an alert would be: 
    - First, verify the actual state of the metrics
    - Identify possible problems and bottlenecks based on the metrics
    - Verify with the team if the possible problems are an unexpected behavior
    - Fix the problems in the application
    - Define with the team thresholds to the alerts
  Usually, a good way to define the threshold is based on the previous records. For example, if the error rate was around 25% in the last month, a good approach would be to define an alert to the error rate above 25%.
  I consider critical conditions error rate next to 100% in any endpoint, traffic next to zero in endpoints with many access, and high latency in endpoints with few data volume. After analyzing the application behavior is easier identify critical conditions. 
