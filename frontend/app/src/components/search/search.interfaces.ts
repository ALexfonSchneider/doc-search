import { Tag } from "react-tagcloud"

export interface SearchProps extends React.InputHTMLAttributes<HTMLInputElement> {}

export interface Suggestions {
    suggestions: string[]
}

export interface MyTag extends Tag {
}