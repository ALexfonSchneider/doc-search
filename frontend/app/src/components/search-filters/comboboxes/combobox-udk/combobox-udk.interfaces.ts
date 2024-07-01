import { Udk } from "@/lib/reducers/search"
import { Action } from "@/lib/utils"

export interface ComboboxUdkProps {
    selectedUdk: Udk | null

    onSelectUdk: Action<Udk | null>
}

export interface UdkSuggestions {
    suggestions: Udk[]
}