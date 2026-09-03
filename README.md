# 💎 BinaryB3ast (GoForge Engine v7.0)

A high-performance, low-level reverse engineering forge and string/opcode patching framework built natively in Go. Formulated completely out of standard library components, **BinaryB3ast** lets security analysts inspect file layouts and execute precision byte modifications safely across macOS (Mach-O), Linux (ELF), and Windows (PE) platforms without causing file corruption or displacement shifting.

```text
      ::::::::  :::    ::: :::::::::  :::::::::: :::::::::  :::::::::  :::    ::: ::::    ::: :::    ::: 
    :+:    :+: :+:    :+: :+:    :+: :+:        :+:    :+: :+:    :+: :+:    :+: :+:+:   :+: :+:   :+:   
    +:+        +:+    +:+ +:+    +:+ +:+        +:+    +:+ +:+    +:+ +:+    +:+ :+:+:+  +:+ +:+  +:+    
    +#+        +#+    +?+ +#++:++#+  +#++:++#   +#++:++#:  +#++:++#+  +#+    +?+ +#+ +:+ +#+ +#++:++     
    +#+        +#+    +#+ +#+        +#+        +#+    +#+ +#+        +#+    +#+ +#+  +#+#+# +#+  +#+    
    #+#    #+# #+#    #+# #+#        #+#        #+#    #+# #+#        #+#    #+# #+#   #+#+# #+#   #+#   
     ########   ########  ###        ########## ###    ### ###         ########  ###    #### ###    ###  
 ───────────────────────────────────────[ NATIVE GO FORGE ENGINE v7.0 ]───────────────────────────────────────
```

## ⚡ Core Operational Features

*   📊 **Shannon Entropy Matrix:** Measures the exact density randomness profile of any binary to flag hidden encryption layers, obfuscators, or packed segments before analysis.
*   ✏️ **Slack-Space Aware Patching:** Modifies text sequences inline using surrounding data segment padding slots, preserving exact offset footprints.
*   🎯 **Adaptive Segment Sieve:** Automatically applies global data pocket allocations to isolate and embed expanded string variations when replacements exceed original length containers.
*   🔓 **Automated UPX Core Unpacker:** Integrates with low-level decompression frameworks to detect compressed `UPX!` markers and strip optimization skins in-place.
*   🔤 **Unicode UTF-16 Channel Scanner:** Rips wide-character text tracks out of application memories that slip completely past traditional ASCII tools.
*   🧩 **Hex/Instruction Modifier:** Rewrites raw assembly loops and control validation checks using hexadecimal arrays (e.g., swapping a conditional branch check instruction for a bypass `NOP` block).
*   🛡️ **Mach-O Entitlement Splicer:** Harvests native XML verification permission schemas from macOS apps, reapplying original security properties over modified ad-hoc seals to prevent instant kernel crashes (`Killed: 9`).
*   📦 **Forge Backup Factory:** Instantly snapshots and stores timestamped copies of your targets within `forge_backups/` before running mutations.

---

## 🛠️ Build and Installation

Ensure you have Go installed on your local environment. Compilation utilizes a tailored, symbolic-stripping `Makefile` to reduce deployment file overhead weights by roughly **60%**.

### Native Installation
Compile the project natively and install it globally across your terminal environment:
```bash
make build
make install
```
*You can now invoke `goforge <target_file>` directly from any directory.*

### Cross-OS Compilation Matrix
To generate highly compressed standalone production packages for all three primary operating systems simultaneously, run:
```bash
make build-all
```
Outputs are directed into your local `./build/` subdirectory pathing:
*   🍏 **macOS:** Universal Fat Binary (`build/goforge_mac`) — bundles Intel + Apple Silicon layers natively.
*   🐧 **Linux:** x86_64 ELF Binary (`build/goforge_linux_x64`).
*   🪟 **Windows:** x86_64 PE Binary (`build/goforge_win_x64.exe`).

---

## 🕹️ Sample Scenarios

### In-Place Opcode Patching (Option 6)
If a security debugger reveals a conditional check byte sequence `74 0A` (`JE` - Jump if Equal) that you need to bypass entirely, select **Option 6**, enter the target, and replace it with `90 90` (`NOP NOP`). The engine automatically commits the modification, builds a safety snapshot copy, and ad-hoc signs the artifact:

```text
 🔍 Enter target hex opcodes to find (e.g. 74 0A): 740A
 ✏️  Enter new bypass opcodes to inject  (e.g. 90 90): 9090
 
 ◢◤ CORE_BUSY // MODIFYING_RAW_MACHINE_INSTRUCTION_OPCODES
     [██████████████████████████████] 100% ── [  ✓  ] ── [READY_OK]
     
 ⚡ ASSEMBLY OPERATION MODIFIED SUCCESSFULLY: check_patched
 📦 SAFE SNAPSHOT RECOVERY SNAPSHOT HOUSED AT: forge_backups/check.20260902_200512.bak
 [*] macOS Mach-O signature verification missing. Applying ad-hoc seal...
 🛡️  HARDENING VALIDATED: Security signature tags successfully realigned!
```
