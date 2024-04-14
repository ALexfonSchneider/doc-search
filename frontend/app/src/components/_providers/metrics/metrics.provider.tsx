import { createContext } from "react";
import { Metrics } from "./metrics-interfaces";
import { useMetrics } from "./metrics.queries";

let metricsInit: Metrics = {
    word_cloud: [],
    years: []
}

export const MetricsContext = createContext<Metrics>(metricsInit);

export const MetricsProvider = ({ children }: { children: JSX.Element }) => {
    const metrics = useMetrics()

    return (
		<MetricsContext.Provider value={metrics.data} children={children}/>
	);
}