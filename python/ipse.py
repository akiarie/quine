#!/usr/bin/env python3
def recurse(clear, shebang):
    print(shebang)
    print(clear)
    repl = clear.replace("\\", "\\\\")
    print(f"recurse(\"\"\"{repl}\"\"\", \"{shebang}\")")

recurse("""def recurse(clear, shebang):
    print(shebang)
    print(clear)
    repl = clear.replace("\\\\", "\\\\\\\\")
    print(f"recurse(\\"\\"\\"{repl}\\"\\"\\", \\"{shebang}\\")")
""", "#!/usr/bin/env python3")
