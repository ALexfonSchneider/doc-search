import axios from "axios";
import { config } from "@/config";
import { UdkSuggestions } from "./combobox-udk.interfaces";

const getSuggestUdk = (query: string) => 
    axios.get<UdkSuggestions>(`${config.api_url}/suggest/udk`, {params: {query: query}}).then(response => response.data)


export const SuggestionsAPI = {
    getSuggestUdk,
}