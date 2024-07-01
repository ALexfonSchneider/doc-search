import { useAppDispatch, useAppSelector } from "@/lib/hooks";
import { FC } from "react";
import { useDocuments as useDocuments } from "./search-content.queries";
import { Article } from "../article/article";
import { Pagination, PaginationContent, PaginationItem, PaginationLink, PaginationNext, PaginationPrevious } from "@/components/ui/pagination";
import { Label } from "@radix-ui/react-label";
import { addKeyword, setPage } from "@/lib/reducers/search";
import { SearchContentPaginationProps } from "./search-content.interfaces";


const SearchContentPagination: FC<SearchContentPaginationProps> = ({page, size, total_size, onSelect}) => {
    return (
        <div className="grid grid-rows-1">
            <Pagination className="mt-2 select-none">
                <PaginationContent>
                    {page - 1 > 0 ? <PaginationPrevious onClick={() => onSelect(page-1)}>{page - 1}</PaginationPrevious> : undefined}
                    <PaginationItem key={page}>
                        <PaginationLink>
                            {page}
                        </PaginationLink>
                    </PaginationItem>
                    {page + 1 <= Number(total_size / size) + 1 ? <PaginationNext onClick={() => onSelect(page+1)}>{page + 1}</PaginationNext> : undefined}
                </PaginationContent>
            </Pagination>
            <Label className="text-sm mr-4">
                total: {total_size}
            </Label>  
        </div>
    )
}


const SearchContent: FC = () => {
    const dispatch = useAppDispatch()
    const [query, selected_keywords, page, size, selected_year, udk] = useAppSelector(state => [state.search.query, state.search.selected_keywords, state.search.page, state.search.size, 
        state.search.selected_year, state.search.selected_udk])

    const onBadgeClick = (value: string) => {
        dispatch(addKeyword(value))
    }
    
    const documents = useDocuments(query, selected_keywords, selected_year, udk, page, size)

    if(documents.isFetching) {
        return <div>
            Loading
        </div>
    }

    if(documents?.data?.articles.length == 0) {
        return <></>
    }

    const onSetPage = (page: number) => dispatch(setPage(page))

    return (
        <div className="grid grid-cols-1 justify-center mt-6 gap-4 sm:w-[70%] md:[100%] m-auto">
            {documents.data?.articles.map(document => (
                <Article document={document} onBadgeClick={onBadgeClick}/>
            ))}

            <SearchContentPagination onSelect={onSetPage} page={documents.data!.page} total_size={documents.data!.total_size} size={documents.data!.size}/>
        </div>
    )
}

export default SearchContent;