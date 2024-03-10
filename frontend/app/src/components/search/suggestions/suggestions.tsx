import React from "react"
import { SuggestionsItemProps, SuggestionsListProps } from "./suggestions.interfases"
import { cn } from "@/lib/utils"

export const SuggestionsList: React.FC<SuggestionsListProps> = ({ className, suggestions, onSelect, ...props}) => {
    return props.active && suggestions.length > 0 && (
    <div className={cn(`relative w-[100%] left-0 top-1 bg-white border-[1px] rounded-sm`, className)}>
        {
            suggestions.map(suggestion =>
                <SuggestionsItem onClick={onSelect} suggestion={suggestion}/>
            )
        }
    </div>
    )
}


export const SuggestionsItem: React.FC<SuggestionsItemProps> =({ className, suggestion, onClick}) => {
    return (
        <div onClick={onClick} onMouseDown={e => {e.preventDefault();}} className={cn(`flex cursor-default items-center rounded-sm px-2 py-1.5 
        text-sm outline-none aria-selected:bg-accent aria-selected:text-accent-foreground 
        data-[disabled]:pointer-events-none data-[disabled]:opacity-50 hover:bg-slate-100`, className)}>
            {suggestion}
        </div>
    )
}


