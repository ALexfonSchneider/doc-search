import { FC, HTMLAttributes, useState } from "react"
import { SearchedDocument, WordCloudItem } from "./article.interfaces"
import { Card, CardContent, CardDescription, CardFooter, CardHeader } from "../ui/card"
import {
    Drawer,
    DrawerContent,
    DrawerTrigger,
} from "@/components/ui/drawer"
import {
    HoverCard,
    HoverCardContent,
    HoverCardTrigger,
} from "@/components/ui/hover-card"
import { BarChart, Bar, Cell, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer, Rectangle, PieChart, Pie } from 'recharts';
import { cn } from "@/lib/utils"
import { Badge } from "../ui/badge"

export interface WordCloudProps {
    data: WordCloudItem[]
}


// TODO: Вынести в отдельный модуль
export const WordCloudLine: FC<WordCloudProps & HTMLAttributes<HTMLDivElement>> = ({data, ...props}) => {
    return (
        <ResponsiveContainer>
            <BarChart
                className={cn(props.className)}
                data={data}
                margin={{
                    top: 5,
                    right: 30,
                    left: 20,
                    bottom: 5,
                }}
                >
                <CartesianGrid strokeDasharray="3 3" />
                <XAxis dataKey="value" />
                <YAxis dataKey="count" />
                <Tooltip />
                <Bar dataKey="count" fill="#82ca9d" activeBar={<Rectangle fill="gold" stroke="purple" />} />
            </BarChart>
        </ResponsiveContainer>
    )
}


// TODO: Вынести в отдельный модуль
export const WordCloudPie: FC<WordCloudProps & HTMLAttributes<HTMLDivElement>> = ({data}) => {
    const RADIAN = Math.PI / 180;
    const COLORS = ['#0088FE', '#00C49F', '#FFBB28', '#FF8042'];

    return (
        <ResponsiveContainer>
            <PieChart>
                <Pie 
                    data={data} 
                    dataKey="count" 
                    nameKey="value" 
                    cx="50%" 
                    cy="50%" 
                    innerRadius={60} 
                    outerRadius={80} 
                    paddingAngle={5}
                    fill="#82ca9d"         
                    label={
                        ({ cx, cy, midAngle, innerRadius, outerRadius, index }) => {
                            const radius = innerRadius + (outerRadius - innerRadius) * 2.5;
                            const x = cx + radius * Math.cos(-midAngle * RADIAN);
                            const y = cy + radius * Math.sin(-midAngle * RADIAN);
                    
                            return (
                                <text x={x} y={y} fill="black" textAnchor={x > cx ? 'start' : 'end'} dominantBaseline="central">
                                    {data[index].value}
                                </text>
                            );
                        }
                    }
                >
                {data.map((_, index) => (
                    <Cell key={`cell-${index}`} fill={COLORS[index % COLORS.length]} />
                ))}
                </Pie>
            </PieChart>
        </ResponsiveContainer>
    );
}

export interface BudgesGroupProps {
    max: number
    children: JSX.Element[]
}

export const BudgesGroup: FC<BudgesGroupProps> = ({children, max}) => {
    const [badgesAll, setBadgesAll] = useState(false)

    return (
        <div>
            <div>
                {badgesAll ? <>
                    {children}
                    <span className="text-sm ml-2 cursor-pointer" onClick={() => setBadgesAll(false)}>...</span>
                </> : 
                <>
                    {children.slice(0, max)}
                    {children.length > max ? <span className="text-sm ml-2 cursor-pointer" onClick={() => setBadgesAll(true)}>...</span> : undefined}
                </>
                }
            </div>
        </div>
    )
}


export const Article: FC<SearchedDocument> = ({document, onBadgeClick}) => {
    const other = document.metrics.word_cloud.slice(30)
    
    const word_cloud_first = document.metrics.word_cloud.slice(0, 50)
    const word_cloud_top_10 = word_cloud_first.slice(0, 10)


    const [highlightOpened, setHighlightOpened] = useState(false)

    
    const handeToolmenuHighlightShowClick = () => {
        setHighlightOpened(state => !state)
    }

    return (
        <Card className="">
            <CardHeader className="text-left text-lg font-medium">
                <a className="text-sm" href={document.article.link}>{document.article.title}</a>
                <CardDescription className="text-xs font-serif text-[#006621]">
                    { document.article.authors.length == 1 ? 
                        `${document.article.authors[0].name}` : `${document.article.authors.map(author => author.name).join(', ')}`
                    }  
                </CardDescription>
            </CardHeader>
            <CardContent className="text-left text-sm">
                {document.article.anatation}
            </CardContent>
            <CardContent className="text-left">
                <BudgesGroup children={[
                        document.article.udk?.trim() == null ? <></> : <Badge className="cursor-default">УДК {document.article.udk}</Badge>,
                        ...document.article.keywords.slice(0, 5).map(keyword =><Badge onClick={() => onBadgeClick(keyword)} className={`ml-1 bg-gray-600 cursor-pointer`}>{keyword}</Badge>)
                    ]} max={3}/>
            </CardContent>
            <CardFooter className="flex-col items-start pb-0">
                <div>
                    {document.highlight ? <span className="text-xs cursor-pointer" onClick={handeToolmenuHighlightShowClick}>Совпадения</span> : <></>}
                    <Drawer>
                        <DrawerTrigger>
                            <span className="text-xs ml-2">Подробнее</span>
                        </DrawerTrigger>
                        <DrawerContent className="h-[80vh] overflow-y-auto after:w-0">
                            <div className="flex flex-row flex-wrap gap-1 w-[100%]">
                                <Card className="flex-grow-[1]">
                                    <CardHeader>Частота ключевых слов</CardHeader>
                                    <CardContent className="h-[300px] m-5">
                                        <WordCloudLine data={word_cloud_first}/>
                                        <HoverCard>
                                            <HoverCardTrigger>
                                                <span  className="cursor-pointer">
                                                    Другие  
                                                </span>
                                            </HoverCardTrigger>
                                            <HoverCardContent className="overflow-auto h-[200px] z-[100]">
                                            {
                                                other.map(metric => {
                                                    return (
                                                        <Badge>
                                                            {metric.value}: {metric.count}
                                                        </Badge>
                                                    )
                                                })
                                            }
                                            </HoverCardContent>
                                        </HoverCard>
                                    </CardContent>
                                </Card>
                                <Card className="basis-[500px]">
                                    <CardHeader>10 самых популярных ключевых слов</CardHeader>
                                    <CardContent className="h-[400px]">
                                        <WordCloudPie data={word_cloud_top_10}/>
                                    </CardContent>
                                </Card>
                            </div>
                        </DrawerContent>
                    </Drawer>
                </div>
                <div className="mt-[10px] pb-4">
                    {highlightOpened &&
                    (<div className="text-left">
                        {document.highlight["article.content"].map(highlight => <li key={highlight} dangerouslySetInnerHTML={{__html: highlight}}/>)}
                    </div>)
                    }
                </div>
            </CardFooter>
        </Card>
    )
}
