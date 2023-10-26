open System.Collections.Generic

type Point2D =
   val mutable x : int
   val mutable y : int
   new (x1, y1) = {x = x1; y = y1;}   

let get_move(m: string): Point2D =
    let s = m.Split(" ")
    let v = s[1] |> int
    match s[0] with
    | "R" -> new Point2D(v, 0)
    | "L" -> new Point2D(-v, 0)
    | "U" -> new Point2D(0, v)
    | "D" -> new Point2D(0, -v)
    | _ -> new Point2D(0, 0)

let calculate_tail(h: Point2D, t: Point2D): Point2D =
    if (h.x = t.x && h.y = t.y) then
        new Point2D(h.x, h.y)
    else if (h.y = t.y && h.x > t.x) then
        new Point2D(h.x-1, t.y)
    else if (h.y = t.y && h.x < t.x) then
        new Point2D(h.x+1, t.y)
    else if (h.x = t.x && h.y > t.y) then 
        new Point2D(h.x, h.y-1)
    else if (h.x = t.x && h.y < t.y) then 
        new Point2D(h.x, h.y+1)  
    else if (abs(h.x - t.x) = 1 && abs(h.y - t.y) = 1) then
        new Point2D(t.x, t.y)
    else if (h.x > t.x && h.y > t.y) then
        new Point2D(t.x+1, t.y+1)
    else if (h.x < t.x && h.y > t.y) then
        new Point2D(t.x-1, t.y+1)
    else if (h.x > t.x && h.y < t.y) then
        new Point2D(t.x+1, t.y-1)
    else if (h.x < t.x && h.y < t.y) then
        new Point2D(t.x-1, t.y-1)
    else
        new Point2D(0, 0)

let solve_rope1(lines: seq<string>) : int =
    let mutable tail = new Point2D(0, 0)
    let mutable head = new Point2D(0, 0)
    let visited = HashSet<(int * int)>()
    visited.Add((0, 0)) |> ignore
    lines |> Seq.iter (fun line -> 
        let m = get_move(line)
        for i in 0..(abs(m.x)) do
            let mutable visiting = true
            match m.x with 
            | x when x > 0 && i <> abs(m.x) -> head.x <- head.x + 1
            | _ when i <> abs(m.x) -> head.x <- head.x - 1
            | _ -> visiting <- false
            match visiting with
            | visiting when visiting ->
                tail <- calculate_tail(head, tail)
                visited.Add((tail.x, tail.y)) |> ignore
            | _ -> ()    

        for j in 0..(abs(m.y)) do
            let mutable visiting = true
            match m.y with
            | y when y > 0 && j <> abs(m.y) -> head.y <- head.y + 1
            | _  when j <> abs(m.y) -> head.y <- head.y - 1
            | _ -> visiting <- false
            match visiting with
            | visiting when visiting ->
                tail <- calculate_tail(head, tail)
                visited.Add((tail.x, tail.y)) |> ignore
            | _ -> ()
    )
    visited.Count

let solve_rope2(lines: seq<string>) : int = 
    let mutable head = new Point2D(0, 0)
    let mutable knots = List[ for i in 1..10 -> new Point2D(0, 0) ]
    let visited = HashSet<(int * int)>()
    visited.Add((0, 0)) |> ignore
    lines |> Seq.iter (fun line ->
        let m = get_move(line)
        for i in 0..(abs(m.x)) do
            let mutable visiting = true
            match m.x with
            | x when x > 0 && i <> abs(m.x) -> head.x <- head.x + 1
            | _ when i <> abs(m.x) -> head.x <- head.x - 1
            | _ -> visiting <- false
            match visiting with
            | visiting when visiting ->
                for t in 0..8 do
                    let mutable previous = head
                    match t with
                    | t when t > 0 -> previous <- knots.Item(t-1)
                    | _ -> ()
                    let res = calculate_tail(previous, knots.Item(t))
                    knots.Item(t).x <- res.x
                    knots.Item(t).y <- res.y
                visited.Add((knots.Item(8).x, knots.Item(8).y)) |> ignore   
            | _ -> () 

        for j in 0..(abs(m.y)) do
            let mutable visiting = true
            match m.y with
            | y when y > 0 && j <> abs(m.y) -> head.y <- head.y + 1
            | _ when j <> abs(m.y) -> head.y <- head.y - 1
            | _ -> visiting <- false
            match visiting with
            | visiting when visiting ->
                for t in 0..8 do
                    let mutable previous = head
                    match t with
                    | t when t > 0 -> previous <- knots.Item(t-1)
                    | _ -> ()
                    let res = calculate_tail(previous, knots.Item(t))
                    knots.Item(t).x <- res.x
                    knots.Item(t).y <- res.y
                visited.Add((knots.Item(8).x, knots.Item(8).y)) |> ignore   
            | _ -> ()
    )
    visited.Count

let lines = seq { yield! System.IO.File.ReadLines "input.txt" }
printfn "part 1: %d" (solve_rope1 lines)
printfn "part 2: %d" (solve_rope2 lines)