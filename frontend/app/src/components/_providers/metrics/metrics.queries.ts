import { useQuery } from '@tanstack/react-query'
import { MetricsAPI } from './metrics.services'

export const useMetrics = () => useQuery({
    queryKey: ["query"],
    retry: false,
    queryFn: () => MetricsAPI.getMetrics(),
    initialData: {
        word_cloud: [],
        years: [],
    }
})