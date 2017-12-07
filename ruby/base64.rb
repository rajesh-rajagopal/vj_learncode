require "base64"

enc   = Base64.encode64('welcome')
puts "Encoded :" 
puts enc
                    # -> "U2VuZCByZWluZm9yY2VtZW50cw==\n"
plain = Base64.decode64(enc)
puts "Decoded :"
puts  plain             
       # -> "Send reinforcements"
