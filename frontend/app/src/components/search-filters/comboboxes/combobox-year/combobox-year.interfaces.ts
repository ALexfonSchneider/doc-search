import { Action } from "@/lib/utils"

export interface ComboboxDateProps {
    selected_year: string | null
    years_available: string[]
    onYearSelect: Action<string | null>
}