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
import { HoverCard, HoverCardTrigger } from "@radix-ui/react-hover-card"
import { FC, useEffect } from "react"
import { ComboboxDateProps } from "./combobox-year.interfaces"


const KeywordsFormSchema = z.object({
    year: z.string().nullable()
})

export const ComboboxYear: FC<ComboboxDateProps> = ({selected_year, years_available, onYearSelect}) => {
    const form = useForm<z.infer<typeof KeywordsFormSchema>>({
        resolver: zodResolver(KeywordsFormSchema),
        defaultValues: {
            year: null
        }
    })

    useEffect(() => {
        form.setValue("year", selected_year)
    }, [selected_year])

    return (
        <Form {...form}>
            <form className="space-y-6">
            <FormField
                control={form.control}
                name="year"
                render={({ field }) => (
                <FormItem className="flex flex-col">
                    <Popover>
                    <PopoverTrigger asChild>
                        <FormControl>
                        <Button
                            variant="outline"
                            role="combobox"
                            className={cn(
                            "w-[100px] p-2 justify-between overflow-hidden text-ellipsis",
                            !field.value && "text-muted-foreground"
                            )}
                        >
                            <HoverCard>
                            <HoverCardTrigger>
                                <span className="pr-2 text-xs">{field.value == null ? "Год не выбран" : field.value}</span> 
                            </HoverCardTrigger>
                            </HoverCard>
                        </Button>
                        </FormControl>
                    </PopoverTrigger>
                    <PopoverContent className="w-[200px] p-0">
                        <Command>
                        <CommandGroup>
                        </CommandGroup>
                            {years_available.map((year) => (
                            <CommandItem
                                value={year}
                                key={year}
                                onSelect={onYearSelect}
                            >
                                {year}
                                <CheckIcon
                                    key={year}
                                    className={cn("ml-auto h-4 w-4", year == selected_year ? "opacity-100" :  "opacity-0")}
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