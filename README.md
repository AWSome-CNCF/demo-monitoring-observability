# Demo Monitoring and Observability

This repository demonstrates the concepts of monitoring, observability, and alerting using modern tools and best practices.

Running containers from a single docker compose configuration file.

Counting ping request from Go application, fire an alert if `ping_request_count > 5`.

## Demo

Grafana:
![Grafana](/images/image.png)

Prometheus:
![Prometheus](/images/image-1.png)

Alert manager:
![Alert manager](/images/image-2.png)

## Features

- **Monitoring**: Collect and visualize metrics from your applications and infrastructure.
- **Observability**: Gain insights into system behavior through logs, metrics, and traces.
- **Alerting**: Configure alerts to notify you of critical issues in real-time.

## Tools Used

- **Prometheus**: For metrics collection and alerting.
- **Grafana**: For data visualization and dashboarding.
- **Alertmanager**: For managing and routing alerts.

## Getting Started

1. Clone the repository:
    ```bash
    git clone https://github.com/AWSome-CNCF/demo-monitoring-observability.git
    cd demo-monitoring-observability
    ```

2. Start the monitoring stack using Docker Compose:
    ```bash
    docker-compose up -d
    ```
    
    2.1. Optional - Restore data from backup
    
    If you need data to play with:
    
    ```bash
    cd data
    docker run --rm -v prom_data:/data -v $(pwd):/backup alpine tar -xzf /backup/prom_data.tar.gz -C /data
    docker run --rm -v grafana_data:/data -v $(pwd):/backup alpine tar -xzf /backup/grafana_data.tar.gz -C /data
    ```


3. Access the following services:
    - Prometheus: [http://localhost:9090](http://localhost:9090)
    - Grafana: [http://localhost:3000](http://localhost:3000)
    - Go: [http://localhost:8080/ping](http://localhost:8080/ping) | [http://localhost:8080/metrics](http://localhost:8080/metrics)
    - Alert manager: [http://localhost:9093](http://localhost:9093)
    - Webhook site: [https://webhook.site/#!/view/317c8cf3-23bd-450e-970b-156ba234104c](https://webhook.site/#!/view/317c8cf3-23bd-450e-970b-156ba234104c)


4. Configure Grafana dashboards and Prometheus alert rules as needed.

## License

This project is licensed under the MIT License.
