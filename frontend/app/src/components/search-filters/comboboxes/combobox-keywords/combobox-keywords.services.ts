import axios from "axios";
import { config } from "@/config";
import { KeywordsSuggestions } from "./combobox-keywords.interfaces";


const getKeywordsSuggestions = (query: string) => 
    axios.get<KeywordsSuggestions>(`${config.api_url}/suggest/keywords`, {params: {query: query}}).then(response => response.data)


export const SuggestionsAPI = {
    getKeywordsSuggestions,
}