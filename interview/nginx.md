# Question
#####1. Some reasons why you should consider using Nginx
- Better serving static content: It can serve them lightning-fast, with minimal processor load. It also comes with an inbuilt gzip compressor.
- Low memory and resource consumption: This can handle many concurrent user requests while consuming a very low memory of around 2.5MB. This enables better utilization of the OS and hardware resources.
- Caching: This provides a cache for both static and dynamic content.
- Act as a Reverse proxy: It forwards the client requests to an appropriate server and vice versa.
- Act as a Load Balancer: This help to allocate resources and dispatch traffic across multiple servers, allowing to scale out very efficiently. Also, it can configure to run multiple backends on the same server.
- Error handling: The application can display proper error pages if it crashes.
- Implement SSL/TLS: This allows for configuring secure HTTPS protocol with SSL. So the information site users receive and send will be protected. Encryption and compression are two heavy computed processes that Node.js is not good at.
- Security: It can prevent DDoS attacks, IP blacklisting and whitelisting, etc