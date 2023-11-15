import json

import behave2cucumber
from behave.__main__ import main as behave_main

code = behave_main()
cucumber_json = behave2cucumber.convert(
    json.load(open("target/cucumber-reports/behave.json"))
)
with open("target/cucumber-reports/Cucumber-behave.json", "w") as report:
    report.write(json.dumps(cucumber_json))
    report.flush()
    report.close()

exit(code)
