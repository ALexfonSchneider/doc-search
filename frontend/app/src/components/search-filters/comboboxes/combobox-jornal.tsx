"use client"

import { zodResolver } from "@hookform/resolvers/zod"
import { CheckIcon } from "@radix-ui/react-icons"
import { useForm } from "react-hook-form"
import * as z from "zod"

import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import {
  Command,
  CommandEmpty,
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
import { useAppDispatch, useAppSelector } from "@/lib/hooks"
import { addJornal, removeJornal } from "@/lib/reducers/search"
import { Input } from "@/components/ui/input"
import { HoverCard, HoverCardTrigger } from "@radix-ui/react-hover-card"

const JornalsFormSchema = z.object({
  jornals: z.string().array()
})

const jornals = [
    { label: "Все о природе", value: "en" },
    { label: "Математика и точные науки", value: "fr" },
    { label: "Физика", value: "de" }
  ] as const

export function ComboboxJornal() {
    const dispatch = useAppDispatch()
    const selected_jornals = useAppSelector(state => state.search.selected_jornals)

    const form = useForm<z.infer<typeof JornalsFormSchema>>({
        resolver: zodResolver(JornalsFormSchema),
        defaultValues: {
            jornals: selected_jornals.map(jornal => jornal.value)
        }
    })

    return (
        <Form {...form}>
            <form className="space-y-6">
            <FormField
                control={form.control}
                name="jornals"
                render={({ field }) => (
                <FormItem className="flex flex-col">
                    <Popover>
                    <PopoverTrigger asChild>
                        <FormControl>
                            <Button
                                variant="outline"
                                role="combobox"
                                className={cn(
                                "w-[200px] justify-between overflow-hidden text-ellipsis",
                                !field.value && "text-muted-foreground"
                                )}
                            >
                                <HoverCard>
                                    <HoverCardTrigger>
                                    <span className="pr-2 text-xs">
                                        {field.value.length > 0
                                        ? jornals.filter(jornal => field.value.includes(jornal.value)).map(jornal => jornal.label).join(", ")
                                        : "Все"}
                                    </span> 
                                    </HoverCardTrigger>
                                </HoverCard>
                            </Button>
                        </FormControl>
                    </PopoverTrigger>
                    <PopoverContent className="w-[200px] p-0">
                        <Command>
                        <Input className="focus-visible:ring-0"></Input>
                        <CommandEmpty>No jornal found.</CommandEmpty>
                        <CommandGroup>
                            {jornals.map((jornal) => (
                            <CommandItem
                                value={jornal.label}
                                key={jornal.label}
                                onSelect={() => {
                                    const index = field.value.indexOf(jornal.value, 0)
                                    let value = field.value;
                                    if (index >= 0) {
                                        value.splice(index, 1)
                                        dispatch(removeJornal(jornal.value))
                                    }
                                    else {
                                        value.push(jornal.value)
                                        dispatch(addJornal({value: jornal.value, label: jornal.label}))
                                    }
                                }}
                            >
                                {jornal.label}
                                <CheckIcon
                                    key={jornal.label}
                                    className={cn(
                                        "ml-auto h-4 w-4",
                                        form.getValues("jornals").includes(jornal.value)
                                        ? "opacity-100"
                                        : "opacity-0"
                                )}
                                />
                            </CommandItem>
                            ))}
                        </CommandGroup>
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
