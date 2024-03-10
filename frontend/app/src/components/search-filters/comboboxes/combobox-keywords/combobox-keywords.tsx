"use client"

import { zodResolver } from "@hookform/resolvers/zod"
import { CheckIcon } from "@radix-ui/react-icons"
import { useForm } from "react-hook-form"
import * as z from "zod"

import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import {
  Command,
  CommandGroup,
  CommandItem,
} from "@/components/ui/command"
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormMessage,
} from "@/components/ui/form"
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover"
import { Input } from "@/components/ui/input"
import { HoverCard, HoverCardTrigger } from "@radix-ui/react-hover-card"
import { FC, useMemo, useState } from "react"
import { ComboboxKeywordsProps } from "./combobox-keywords.interfaces"
import { useKeywordsSuggestions } from "./combobox-keyword.queries"


const KeywordsFormSchema = z.object({
    keywords: z.string().array()
})

export const ComboboxKeywords: FC<ComboboxKeywordsProps> = ({selected_keywords, onAddKeyword, onRemoveKeyword}) => {
    const [query, setQuery] = useState("")
    const {data: keywords} = useKeywordsSuggestions(query)

    const form = useForm<z.infer<typeof KeywordsFormSchema>>({
        resolver: zodResolver(KeywordsFormSchema),
        defaultValues: {
            keywords: selected_keywords
        }
    })

    useMemo(() => {
        form.setValue("keywords", selected_keywords)
    }, [selected_keywords])

    const onPooverBlur = () => setQuery("")

    return (
        <Form {...form}>
            <form className="space-y-6">
            <FormField
                control={form.control}
                name="keywords"
                render={({ field }) => (
                <FormItem className="flex flex-col">
                    <Popover>
                    <PopoverTrigger asChild onBlur={onPooverBlur}>
                        <FormControl>
                        <Button
                            variant="outline"
                            role="combobox"
                            className={cn(
                            "w-[200px] p-2 justify-between overflow-hidden text-ellipsis",
                            !field.value && "text-muted-foreground"
                            )}
                        >
                            <HoverCard>
                            <HoverCardTrigger>
                                {field.value.length > 0
                                ? selected_keywords.join(", ")
                                : <span className="pr-2 text-xs">Ключевые слова не выбранны</span>} 
                            </HoverCardTrigger>
                            </HoverCard>
                        </Button>
                        </FormControl>
                    </PopoverTrigger>
                    <PopoverContent className="w-[200px] p-0">
                        <Command>
                            <Input className="focus-visible:ring-0" onChange={(event) => setQuery(event.currentTarget.value)}></Input>
                        <CommandGroup>
                        </CommandGroup>
                            {(query ? keywords : selected_keywords).map((keyword) => (
                            <CommandItem
                                value={keyword}
                                key={keyword}
                                onSelect={() => {
                                    if(selected_keywords.includes(keyword)) {
                                        onRemoveKeyword(keyword)
                                    }
                                    else {
                                        onAddKeyword(keyword)
                                    }
                                }}
                            >
                                {keyword}
                                <CheckIcon
                                    key={keyword}
                                    className={cn(
                                        "ml-auto h-4 w-4",
                                        selected_keywords.includes(keyword)
                                        ? "opacity-100"
                                        : "opacity-0"
                                )}
                                />
                            </CommandItem>
                            ))}
                        </Command>
                    </PopoverContent>
                    </Popover>
                    <FormMessage />
                </FormItem>
                )}
            />
            </form>
        </Form>
    )
}