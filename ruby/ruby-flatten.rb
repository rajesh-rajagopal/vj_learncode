a =  {1=> "one", 2 => [2,"two"], 3 => "three"}
puts a.flatten    # => [1, "one", 2, [2, "two"], 3, "three"]
a.flatten(2) # => [1, "one", 2, 2, "two", 3, "three"]


