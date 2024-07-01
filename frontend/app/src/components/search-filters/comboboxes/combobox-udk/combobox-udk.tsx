"use client"

import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import * as z from "zod"

import { cn } from "@/lib/utils"
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
import { FC, useEffect, useState } from "react"
import { ComboboxUdkProps } from "./combobox-udk.interfaces"
import { useUdkSuggestions } from "./combobox-udk.queries"
import { Command, CommandGroup, CommandItem } from "@/components/ui/command"
import { Input } from "@/components/ui/input"
import { HoverCard, HoverCardTrigger } from "@/components/ui/hover-card"
import { Button } from "@/components/ui/button"
import { CheckIcon } from "@radix-ui/react-icons"


const FormSchema = z.object({
    udk: z.object({
        code: z.string(),
        name: z.string(),
        query: z.string()
    }).nullable()
})

export const ComboboxUdk: FC<ComboboxUdkProps> = ({selectedUdk, onSelectUdk}) => {
    const [query, setQuery] = useState("")
    const {data: suggested_udks} = useUdkSuggestions(query)

    const form = useForm<z.infer<typeof FormSchema>>({
        resolver: zodResolver(FormSchema),
        defaultValues: {
            udk: null
        }
    })

    useEffect(() => {
        if (selectedUdk) {
            form.setValue("udk", selectedUdk!)
        }
    })

    return (
        <Form {...form}>
            <form className="space-y-6">
            <FormField
                control={form.control}
                name="udk"
                render={({ field }) => (
                <FormItem className="flex flex-col">
                    <Popover>
                    <PopoverTrigger asChild>
                        <FormControl>
                        <Button
                            variant="outline"
                            role="combobox"
                            className={cn(
                            "p-2 w-[250px] justify-between overflow-hidden text-ellipsis",
                            !field.value && "text-muted-foreground"
                            )}
                        >
                            <HoverCard>
                            <HoverCardTrigger>
                                <span className="pr-2 text-xs">{field.value == null ? "УДК не выбран" : field.value.query}</span> 
                            </HoverCardTrigger>
                            </HoverCard>
                        </Button>
                        </FormControl>
                    </PopoverTrigger>
                    <PopoverContent className="w-[200px] p-0">
                        <Command>
                            <Input value={query} className="focus-visible:ring-0" onChange={(event) => {
                                setQuery(event.currentTarget.value)
                            }}></Input>
                        <CommandGroup>
                        </CommandGroup>
                            {suggested_udks.map((udk) => (
                            <CommandItem
                                value={udk.code}
                                key={udk.code}
                                onSelect={() => {
                                    if (udk.code == selectedUdk?.code) {
                                        setQuery("")   
                                        onSelectUdk(null)
                                        form.setValue("udk", null)
                                        return
                                    }
                                    setQuery(udk.query)
                                    onSelectUdk(udk)
                                }}
                            >
                                {udk.query}
                                <CheckIcon
                                    key={udk.code}
                                    className={cn("ml-auto h-4 w-4", udk.code == selectedUdk?.code ? "opacity-100" :  "opacity-0")}
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