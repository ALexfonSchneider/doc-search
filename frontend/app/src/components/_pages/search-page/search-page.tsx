
import { FC } from "react";
import { Search } from "@/components/search/search";
import { TagCloud, TagEventHandler } from "react-tagcloud";
import { useTags } from "@/components/search/search.queries";
import { useAppDispatch } from "@/lib/hooks";
import { setQuery } from "@/lib/reducers/search";
import SearchContent from "@/components/search-content/search-content";

export const SearchPage: FC = () => {
    const dispatch = useAppDispatch()

    const onCloudWordClick: TagEventHandler = (tag, _) => {
        dispatch(setQuery(tag.value))
    }

    const tags = useTags()

    return (
        <div>
            <TagCloud className="select-none cursor-pointer" onClick={onCloudWordClick} minSize={12} maxSize={32} tags={tags.data || []}/>

            <Search className="m-auto mt-2 w-2/3"/>

            <SearchContent/>
        </div>
    )
}