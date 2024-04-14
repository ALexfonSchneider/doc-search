import { Tag } from "react-tagcloud"

export interface DocumentInYears {
    year: string
    count: number
}


export interface Metrics {
    word_cloud: Tag[]
    years: DocumentInYears[]
}