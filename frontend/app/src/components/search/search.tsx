
import { Input } from "@/components/ui/input";
import { FC, useMemo, useState } from "react";
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
import { addKeyword, removeKeyword, setQuery } from "@/lib/reducers/search";
import { Badge } from "../ui/badge";
import { Cross1Icon } from "@radix-ui/react-icons";
import { ComboboxJornal } from "../search-filters/comboboxes/combobox-jornal";
import { ComboboxKeywords } from "../search-filters/comboboxes/combobox-keywords/combobox-keywords";

export const Search: FC<SearchProps> = ({className}) => {
    const dispatch = useAppDispatch()

    const [query, keywords] = useAppSelector(state => [state.search.query, state.search.selected_keywords])

    const [localQuery, setLocalQuery] = useState<string>("")
    const [debauncesQuery] = useDebounce(localQuery, 500)

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
                            <div className="flex flex-row gap-2">
                                <ComboboxJornal/>
                                <ComboboxKeywords selected_keywords={keywords} onRemoveKeyword={onRemoveKeyword} onAddKeyword={onAddKeyword}/>
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