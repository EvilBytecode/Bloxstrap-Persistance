# Bloxstrap-Persistance

Bloxstrap-Persistance is a proof-of-concept demonstration showcasing how to abuse the configuration files of a hypothetical application named Bloxstrap. By persistently modifying its settings file (`Settings.json`), this project adds custom integrations, potentially compromising the integrity and functionality of Bloxstrap.

## Proof of Concept (PoC)

Watch the PoC video to see the demonstration in action:
- [View PoC Video](https://streamable.com/fi1qp7)


### Prerequisites

- Go (Golang) environment to compile and run the code.
- Bloxstrap installed on the target system.

### Execution

1. **Locate Settings File:**
   - The program attempts to find the `Settings.json` file in the user's Bloxstrap directory (`AppData/Local/Bloxstrap`).

2. **Modify Settings:**
   - If the `Settings.json` file exists, the program reads its current settings into a data structure.
   - It then appends a custom integration (`newmalinter`) with predefined values:
     - **Name:** "Evilbytecode was here"
     - **Location:** "C:\\Windows\\System32\\cmd.exe"
     - **LaunchArgs:** "start cmd.exe"
     - **AutoClose:** false

3. **Persist Changes:**
   - The modified settings, now including the malicious integration, are written back to `Settings.json`, ensuring persistence across application launches.

### Limitations

- **Dependency on Bloxstrap:** The success of this demonstration relies on Bloxstrap being installed and the `Settings.json` file being present in the specified directory.
- **Ethical Considerations:** Modifying software settings without consent can be unethical and potentially illegal. This code is strictly for educational purposes.

### Disclaimer
- This project is intended for educational purposes only. Modifying software without explicit permission may violate terms of service and laws in your jurisdiction. Use responsibly and with caution.


## License
This project is licensed under the MIT License. See the LICENSE file for details.