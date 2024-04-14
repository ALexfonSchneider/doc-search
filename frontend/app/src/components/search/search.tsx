
import { Input } from "@/components/ui/input";
import { FC, useContext, useMemo, useState } from "react";
import React from "react";
import { SuggestionsList } from "./suggestions/suggestions";
import { SearchProps } from "./search.interfaces";
import { cn } from "@/lib/utils";
import { useDebounce } from "use-debounce";
import { useQuerySuggestions } from "./search.queries";
import {
    Accordion,
    AccordionContent,
    AccordionItem,
    AccordionTrigger,
} from "@/components/ui/accordion"

import { useAppDispatch, useAppSelector } from "@/lib/hooks";
import { addKeyword, removeKeyword, setQuery, setYear } from "@/lib/reducers/search";
import { Badge } from "../ui/badge";
import { Cross1Icon } from "@radix-ui/react-icons";
import { ComboboxJornal } from "../search-filters/comboboxes/combobox-jornal";
import { ComboboxKeywords } from "../search-filters/comboboxes/combobox-keywords/combobox-keywords";
import { ComboboxYear } from "../search-filters/comboboxes/combobox-year/combobox-year";
import { MetricsContext } from "../_providers/metrics/metrics.provider";

export const Search: FC<SearchProps> = ({className}) => {
    const dispatch = useAppDispatch()

    const [query, keywords, selected_year] = useAppSelector(state => [state.search.query, state.search.selected_keywords, state.search.selected_year])

    const [localQuery, setLocalQuery] = useState<string>("")
    const [debauncesQuery] = useDebounce(localQuery, 500)

    const years = useContext(MetricsContext)

    useMemo(() => {
        setLocalQuery(query)
    }, [query])
    
    const suggestions = useQuerySuggestions(debauncesQuery)
    const [suggestionActive, setSuggestionActive] = useState<boolean>(false)

    const OnSearch = (query: string) => {
        setSuggestionActive(false)
        dispatch(setQuery(query))
    }

    const onQueryChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setLocalQuery(event.currentTarget.value)
    }

    const keyDownHandler = (event: React.KeyboardEvent<HTMLInputElement>) => {
        if (event.code === "Enter") {
            OnSearch(event.currentTarget.value)
        }
    };

    const onSuggestionSelect = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.stopPropagation()
        OnSearch(event.currentTarget.textContent || "")
    };

    const onSearchLineClick = () => {
        setSuggestionActive(true)
    };

    const onSearchLineBlur = () => {
        setSuggestionActive(false)
    };

    const onYearSelect = (value: string | null) => {
        if (selected_year == value) {
            dispatch(setYear(null))
        }
        else {
            dispatch(setYear(value))
        }
    }

    const onAddKeyword = (keyword: string) => dispatch(addKeyword(keyword))
    const onRemoveKeyword = (keyword: string) => dispatch(removeKeyword(keyword))

    return (
        <div>
            <div className={cn("relative m-auto w-2/3", className)}>
                <Input placeholder="input search query" type="search" autoFocus={false} value={localQuery} onKeyDown={keyDownHandler} onChange={onQueryChange} onClick={onSearchLineClick} onBlur={onSearchLineBlur}/>
                <div className="absolute z-10 w-[100%]">
                    <SuggestionsList onSelect={onSuggestionSelect} active={suggestionActive} suggestions={suggestions.data}/>
                </div>
                <Accordion type="multiple" defaultValue={[]}>
                    <AccordionItem value="filters">
                        <AccordionTrigger className="text-sm select-none">Фильтры</AccordionTrigger>
                        <AccordionContent>
                            <div className="flex flex-row flex-wrap md:justify-start sl:justify-center gap-2">
                                <ComboboxJornal/>
                                <ComboboxKeywords selected_keywords={keywords} onRemoveKeyword={onRemoveKeyword} onAddKeyword={onAddKeyword}/>
                                <ComboboxYear selected_year={selected_year} years_available={years.years.map(d => d.year)} onYearSelect={onYearSelect}/>
                            </div>
                        </AccordionContent>
                    </AccordionItem>
                </Accordion>
                <div>
                    {keywords.map(keyword => <Badge className="ml-2">
                        {keyword} <Cross1Icon onClick={() => onRemoveKeyword(keyword)} className="pl-1" />
                    </Badge>)}
                </div>
            </div>
        </div>
    )
}