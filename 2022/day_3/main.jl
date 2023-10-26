function countmap(s::String)
    res::Dict{Char, Int} = Dict()
    foreach(c -> res[only(c)] = get(res, c, 0) + 1, split(s, ""))
    return res
end

function findcommon(dict1::Dict, dict2::Dict, dict3=nothing)
    for (k, _) in dict1
        if isnothing(dict3) && haskey(dict2, k)
            return k
        elseif !isnothing(dict3) && haskey(dict2, k) && haskey(dict3, k)
            return k 
        end
    end
    return ""
end

function prioritymap()
    res::Dict{Char, Int}, lower_alphabet = Dict(), "abcdefghijklmnopqrstuvwxyz"
    alphabet = lower_alphabet * uppercase(lower_alphabet)
    foreach((i, c) -> res[only(c)] = i, 1:length(alphabet), split(alphabet, ""))
    return res
end

first_sum, second_sum = open("input.txt") do io
    item_priority, first_sum, second_sum, lines = prioritymap(), 0, 0, readlines(io)
    for l in lines    
       f_dict = countmap(String(SubString(l, 1, Int(length(l)/2))))
       s_dict = countmap(String(SubString(l, Int(length(l)/2) + 1, length(l))))
       item = findcommon(f_dict, s_dict)
       first_sum += get(item_priority, item, 0)
    end

    for i = 1:3:length(lines)
        f_dict = countmap(lines[i])
        s_dict = countmap(lines[i+1])
        t_dict = countmap(lines[i+2])
        item = findcommon(f_dict, s_dict, t_dict)
        second_sum += get(item_priority, item, 0)
    end
    (first_sum, second_sum)
 end
println("round1:", first_sum)
println("round2:", second_sum)