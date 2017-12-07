str = "Hey what's up @dude, @how's it going?"
h = str.gsub!(/.*?(?=@how)/im, "") #=> "@how's it going?"
def hello(l)
 puts l 
end

hello(h)
