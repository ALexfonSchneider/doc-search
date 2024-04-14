import axios from "axios"
import { Metrics } from "./metrics-interfaces"
import { config } from "@/config"

const getMetrics = () => 
    axios.get<Metrics>(
        `${config.api_url}/metrics`, 
        {
            responseType: "json",
        }).then(response => {
            return response.data
        })

export const MetricsAPI = {
    getMetrics
}
    