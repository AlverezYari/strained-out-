export interface Position {
  x: number
  y: number
}

export interface Tile {
  id: string
  position: Position
  image: string
}

export interface Unit {
  id: string
  name: string
  health: number
  position: Position
  conditions: string[]
  actions: number
  isEnemy: boolean
}

export interface Trigger {
  id: string
  round?: number
  location?: Position
  condition?: string
  action: string
}

export interface MapState {
  width: number
  height: number
  tiles: Tile[]
}

export interface Mission {
  id: string
  name: string
  map: MapState
  units: Unit[]
  triggers: Trigger[]
  round: number
}

export interface State {
  mission: Mission
}
