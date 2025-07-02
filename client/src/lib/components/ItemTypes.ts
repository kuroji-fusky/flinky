export interface ItemProps {
  title: string
  img?: string
  meta?: Partial<{
    species: string
    franchise: string
  }>
}