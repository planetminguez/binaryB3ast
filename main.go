

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// ANSI Cyberpunk Color Matrix
const (
	Green   = "\033[38;5;82m"
	Red     = "\033[38;5;196m"
	Yellow  = "\033[38;5;226m"
	Blue    = "\033[38;5;27m"
	Magenta = "\033[38;5;201m"
	Cyan    = "\033[38;5;51m"
	Orange  = "\033[38;5;208m"
	White   = "\033[38;5;255m"
	Gray    = "\033[38;5;242m"
	Reset   = "\033[0m"
)

func runLoadingBar(taskLabel string) {
	barWidth := 30
	pulseChars := []string{"►", "»", "█", "█"}
	spinnerChars := []string{"/", "—", "\\", "|"}
	tickStates := []string{"MEM_ALLOC", "SEC_SCAN", "BIT_FLIP", "SIG_ALIGN", "IO_FLUSH"}
	
	fmt.Printf(" %s◢◤ CORE_BUSY //%s %s%s%s\n", Cyan, Reset, White, taskLabel, Reset)
	
	for i := 0; i <= barWidth; i++ {
		pct := (i * 100) / barWidth
		pulseChar := pulseChars[i%4]
		spinner := spinnerChars[i%4]
		tickTag := tickStates[i%5]
		
		filler := strings.Repeat(pulseChar, i)
		emptySlots := strings.Repeat("·", barWidth-i)
		
		spectralColor := Red
		if pct > 35 {
			spectralColor = Orange
		}
		if pct > 70 {
			spectralColor = Green
		}
		
		fmt.Printf("\r     %s[%s%s%s%s] %s%3d%%%s ── %s[ %s ]%s ── %s[%s%s%s]%s",
			Gray, spectralColor, filler, Gray, emptySlots, White, pct, Reset, Cyan, spinner, Reset, Gray, Magenta, tickTag, Gray, Reset)
		
		time.Sleep(20 * time.Millisecond)
	}
	fmt.Printf("\r     %s[%s] %s100%%%s ── %s[  ✓  ]%s ── %s[READY_OK]%s\n\n",
		Gray, strings.Repeat("█", barWidth), Green, Reset, Green, Reset, Green, Reset)
}

func generateBackup(srcPath string) (string, error) {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return "", err
	}
	defer srcFile.Close()
	
	backupDir := "forge_backups"
	_ = os.Mkdir(backupDir, 0755)
	
	timestamp := time.Now().Format("20060102_150405")
	baseName := filepath.Base(srcPath)
	destPath := filepath.Join(backupDir, fmt.Sprintf("%s.%s.bak", baseName, timestamp))
	
	destFile, err := os.Create(destPath)
	if err != nil {
		return "", err
	}
	defer destFile.Close()
	
	_, err = io.Copy(destFile, srcFile)
	return destPath, err
}

func calculateEntropy(data []byte) float64 {
	if len(data) == 0 {
		return 0
	}
	counts := make(map[byte]int)
	for _, b := range data {
		counts[b]++
	}
	var entropy float64
	for _, count := range counts {
		px := float64(count) / float64(len(data))
		entropy -= px * math.Log2(px)
	}
	return entropy
}

func getTargetMetadata(filename string, size int64) string {
	f, err := os.Open(filename)
	if err != nil {
		return "UNKNOWN_OBJECT"
	}
	defer f.Close()
	
	header := make([]byte, 4)
	_, _ = f.Read(header)
	
	if bytes.Equal(header, []byte{0xCF, 0xFA, 0xED, 0xFE}) || bytes.Equal(header, []byte{0xFE, 0xED, 0xFA, 0xCF}) {
		return "Mach-O 64-bit Executable (macOS)"
	}
	if bytes.Equal(header, []byte{0x7F, 'E', 'L', 'F'}) {
		return "ELF Executable Object (Linux)"
	}
	if bytes.Equal(header[:2], []byte{'M', 'Z'}) {
		return "PE Portable Executable (Windows)"
	}
	return "Generic Binary Payload"
}

func showMenu(filename string, metadata string, size int64) {
	fmt.Print("\033[H\033[2J")
	fmt.Printf("%s      ::::::::  :::    ::: :::::::::  :::::::::: :::::::::  :::::::::  :::    ::: ::::    ::: :::    ::: %s\n", Cyan, Reset)
	fmt.Printf("%s    :+:    :+: :+:    :+: :+:    :+: :+:        :+:    :+: :+:    :+: :+:    :+: :+:+:   :+: :+:   :+:   %s\n", Cyan, Reset)
	fmt.Printf("%s ───────────────────────────────────────[ NATIVE GO FORGE ENGINE v6.0 ]───────────────────────────────────────%s\n", Gray, Reset)
	fmt.Printf("%s ┌─[ %sTARGET OBJECT METADATA%s ]────────────────────────────────────────────────────────────────────────────%s\n", Gray, White, Gray, Reset)
	fmt.Printf("%s │ %sFILE:%s %s\n", Gray, Yellow, Reset, filename)
	fmt.Printf("%s │ %sSIZE:%s %d bytes\n", Gray, Yellow, Reset, size)
	fmt.Printf("%s │ %sARCH:%s %s\n", Gray, Yellow, Reset, metadata)
	fmt.Printf("%s └───────────────────────────────────────────────────────────────────────────────────────────────────────%s\n\n", Gray, Reset)
	fmt.Printf("   %s◢◤ DYNAMIC SYSTEM MODULES%s\n", Cyan, Reset)
	fmt.Printf("   %s│%s  %s1)%s Scan Shannon Data Entropy Matrix\n", Gray, Reset, Green, Reset)
	fmt.Printf("   %s│%s  %s2)%s Execute Slack-Space Aware String Patch\n", Gray, Reset, Green, Reset)
	fmt.Printf("   %s│%s  %s3)%s Precision Placement Tool (Confirmation Interface)\n", Gray, Reset, Green, Reset)
	fmt.Printf("   %s│%s  %s4)%s Deploy Automated Decompression Matrix (UPX Core Unpacker)\n", Gray, Reset, Orange, Reset)
	fmt.Printf("   %s│%s  %s5)%s Extract Wide-Character (UTF-16/Unicode) Hidden Strings\n", Gray, Reset, Cyan, Reset)
	fmt.Printf("   %s│%s  %s6)%s Raw Hex/Instruction Patching Panel (JE -> NOP, etc.)\n", Gray, Reset, Magenta, Reset)
	fmt.Printf("   %s│%s  %s7)%s Mach-O XML Entitlement Splicer (macOS Hardening Seal)\n", Gray, Reset, Orange, Reset)
	fmt.Printf("   %s│%s  %s8)%s Abort Terminal Session\n", Gray, Reset, Red, Reset)
	fmt.Printf("%s ───────────────────────────────────────────────────────────────────────────────────────────────────────%s\n", Gray, Reset)
	fmt.Printf(" %sforge@operator:%s~$%s ", White, Cyan, Reset)
}



func main() {	
	if len(os.Args) < 2 {
		fmt.Printf("%sError:%s Missing target binary argument.\nUsage: go run main.go <binary_file>\n", Red, Reset)
		os.Exit(1)
	}
	
	targetFile := os.Args[1]
	fileInfo, err := os.Stat(targetFile)
	if err != nil {
		fmt.Printf("%sError:%s Target file '%s' not accessible.\n", Red, Reset, targetFile)
		os.Exit(1)
	}
	
	metadata := getTargetMetadata(targetFile, fileInfo.Size())
	
	for {
		showMenu(targetFile, metadata, fileInfo.Size())
		var choice int
		_, err := fmt.Scanf("%d\n", &choice)
		if err != nil {
			var discard string
			_, _ = fmt.Scanf("%s\n", &discard)
			continue
		}
		
		switch choice {
		case 1:
			fmt.Print("\033[H\033[2J")
			fmt.Printf("%s>>> PARSING FILE ENTROPY PROFILE...%s\n\n", Blue, Reset)
			runLoadingBar("ANALYZE_BYTE_DENSITY_STREAMS")
			
			data, err := os.ReadFile(targetFile)
			if err != nil {
				fmt.Printf(" %s❌ Error reading file memory.%s\n", Red, Reset)
			} else {
				entropy := calculateEntropy(data)
				fmt.Printf(" 💾 Absolute Entropy Score: %.4f / 8.0000\n", entropy)
				if entropy > 7.2 {
					fmt.Printf("\n%s ☢️  CRITICAL: Target data structure highly randomized. PACKED or ENCRYPTED.%s\n", Red, Reset)
				} else {
					fmt.Printf("\n%s ✅ NOMINAL: Exposed bytecode detected. File structure is safe to patch.%s\n", Green, Reset)
				}
			}
			fmt.Printf("\n%sPress [ENTER] to return to menu...%s", Gray, Reset)
			_, _ = fmt.Scanln()
			
		case 2:
			fmt.Print("\033[H\033[2J")
			fmt.Printf("%s>>> ADAPTIVE BOUNDARY PATCH CONTROL PANEL%s\n\n", Magenta, Reset)
			
			var findStr, replaceStr string
			fmt.Printf(" %sTarget Text to Search:%s ", White, Reset)
			_, _ = fmt.Scanf("%s\n", &findStr)
			fmt.Printf(" %sNew Replacement Text :%s ", White, Reset)
			_, _ = fmt.Scanf("%s\n", &replaceStr)
			
			if len(findStr) == 0 || len(replaceStr) == 0 {
				fmt.Printf("\n%s ❌ Error: Query fields cannot be empty.%s\n", Red, Reset)
				fmt.Printf("\n%sPress [ENTER] to return...%s", Gray, Reset)
				_, _ = fmt.Scanln()
				continue
			}
			
			data, err := os.ReadFile(targetFile)
			if err != nil {
				fmt.Printf(" %s❌ Error mapping binary memory.%s\n", Red, Reset)
				continue
			}
			
			findBytes := []byte(findStr)
			idx := bytes.Index(data, findBytes)
			if idx == -1 {
				fmt.Printf("\n%s ❌ MALFUNCTION: String signature not found anywhere in workspace.%s\n", Red, Reset)
				fmt.Printf("\n%sPress [ENTER] to return...%s", Gray, Reset)
				_, _ = fmt.Scanln()
				continue
			}
			
			backupPath, backupErr := generateBackup(targetFile)
			if backupErr != nil {
				fmt.Printf("\n%s ⚠️  CRITICAL SAFETY WARNING: Failed to generate recovery snapshot: %s%s\n", Yellow, backupErr.Error(), Reset)
				fmt.Printf("\n%sPress [ENTER] to return...%s", Gray, Reset)
				_, _ = fmt.Scanln()
				continue
			}
			
			slackCount := 0
			for i := idx + len(findBytes); i < len(data); i++ {
				if data[i] == 0x00 {
					slackCount++
				} else {
					break
				}
			}
			availableContainerSize := len(findBytes) + slackCount
			
			var patchedData []byte
			
			if len(replaceStr) > availableContainerSize {
				fmt.Printf("\n%s [!] In-place overflow detected. Deploying Segment Sieve...%s\n", Yellow, Reset)
				paddingBlock := make([]byte, 16)
				slackPaddingIdx := bytes.LastIndex(data, paddingBlock)
				if slackPaddingIdx == -1 {
					fmt.Printf("\n%s ❌ FAIL: Target file has zero global data padding pockets available.%s\n", Red, Reset)
					fmt.Printf("\n%sPress [ENTER] to return...%s", Gray, Reset)
					_, _ = fmt.Scanln()
					continue
				}
				
				patchedData = append([]byte(nil), data...)
				copy(patchedData[slackPaddingIdx:], []byte(replaceStr))
				truncatedReplace := replaceStr[:len(findBytes)]
				copy(patchedData[idx:], []byte(truncatedReplace))
			} else {
				replaceBytes := []byte(replaceStr)
				paddingSize := availableContainerSize - len(replaceBytes)
				padding := make([]byte, paddingSize)
				replaceBytes = append(replaceBytes, padding...)
				targetPatchZone := data[idx : idx+availableContainerSize]
				
				runLoadingBar("COMMITTING_SLACK_SPACE_BYTE_MUTATIONS")
				patchedData = bytes.Replace(data, targetPatchZone, replaceBytes, 1)
			}
			
			localOutputFile := targetFile + "_patched"
			err = os.WriteFile(localOutputFile, patchedData, 0755)
			if err != nil {
				fmt.Printf(" %s❌ Access Denied: Failed to output modified file layer.%s\n", Red, Reset)
			} else {
				fmt.Printf(" %s⚡ SUCCESS: Patched binary saved to: %s%s\n", Green, localOutputFile, Reset)
				fmt.Printf(" %s📦 SAFETY SNAPSHOT WRITTEN TO: %s%s\n", Gray, backupPath, Reset)
			}
			fmt.Printf("\n%sPress [ENTER] to return...%s", Gray, Reset)
			_, _ = fmt.Scanln()
			
		case 3:
			fmt.Print("\033[H\033[2J")
			fmt.Printf("%s>>> INTERACTIVE PRECISION PLACEMENT ENGINE%s\n\n", Cyan, Reset)
			
			var findStr, replaceStr string
			fmt.Printf(" %sEnter target string to examine:%s ", White, Reset)
			_, _ = fmt.Scanf("%s\n", &findStr)
			fmt.Printf(" %sEnter new replacement payload:%s ", White, Reset)
			_, _ = fmt.Scanf("%s\n", &replaceStr)
			
			if len(findStr) == 0 || len(replaceStr) == 0 {
				fmt.Printf("\n%s ❌ Error: Query matrices cannot be blank.%s\n", Red, Reset)
				fmt.Printf("\n%sPress [ENTER] to return...%s", Gray, Reset)
				_, _ = fmt.Scanln()
				continue
			}
			
			data, err := os.ReadFile(targetFile)
			if err != nil {
				fmt.Printf(" %s❌ Error reading workspace binary.%s\n", Red, Reset)
				continue
			}
			
			findBytes := []byte(findStr)
			idx := bytes.Index(data, findBytes)
			if idx == -1 {
				fmt.Printf("\n%s ❌ TARGET REJECTED: String not found anywhere in layout memory.%s\n", Red, Reset)
				fmt.Printf("\n%sPress [ENTER] to return...%s", Gray, Reset)
				_, _ = fmt.Scanln()
				continue
			}
			
			slackCount := 0
			for i := idx + len(findBytes); i < len(data); i++ {
				if data[i] == 0x00 {
					slackCount++
				} else {
					break
				}
			}
			availableContainerSize := len(findBytes) + slackCount
			
			fmt.Printf("\n%s┌─[ LAYOUT FEASIBILITY TELEMETRY REPORT ]────────────────────────────────────────┐%s\n", Gray, Reset)
			fmt.Printf(" │ Absolute Target Offset : %s0x%X%s\n", Green, idx, Reset)
			fmt.Printf(" │ Original Allocation Width: %s%d characters%s\n", Yellow, len(findBytes), Reset)
			fmt.Printf(" │ Available Section Slack : %s%d trailing null bytes%s\n", Yellow, slackCount, Reset)
			fmt.Printf(" │ Combined Safe Max Limit : %s%d characters%s\n", Cyan, availableContainerSize, Reset)
			fmt.Printf(" │ New Replacement Size    : %s%d characters%s\n", White, len(replaceStr), Reset)
			
			var strategy string
			if len(replaceStr) > availableContainerSize {
				strategy = fmt.Sprintf("%sStrategy B (Global Sieve Injection / Truncation)%s", Orange, Reset)
			} else {
				strategy = fmt.Sprintf("%sStrategy A (Direct In-Place Slack Alignment)%s", Green, Reset)
			}
			fmt.Printf(" │ Resolution Strategy    : %s\n", strategy)
			fmt.Printf("%s└────────────────────────────────────────────────────────────────────────────────┘%s\n\n", Gray, Reset)
			
			var confirm string
			fmt.Printf(" %s⚠️  CRITICAL COMMIT: Apply this bytecode patch to '%s'? [Y/N]:%s ", Yellow, targetFile, Reset)
			_, _ = fmt.Scanf("%s\n", &confirm)
			
			if strings.ToUpper(confirm) != "Y" {
				fmt.Printf("\n%s 🛑 Transaction aborted by operator. Zero files written.%s\n", Red, Reset)
				fmt.Printf("\n%sPress [ENTER] to return...%s", Gray, Reset)
				_, _ = fmt.Scanln()
				continue
			}
			
			backupPath, backupErr := generateBackup(targetFile)
			if backupErr != nil {
				fmt.Printf("\n%s ❌ Safety halt: Backup engine malfunction: %s%s\n", Red, backupErr.Error(), Reset)
				fmt.Printf("\n%sPress [ENTER] to return...%s", Gray, Reset)
				_, _ = fmt.Scanln()
				continue
			}
			
			var patchedData []byte
			if len(replaceStr) > availableContainerSize {
				paddingBlock := make([]byte, 16)
				slackPaddingIdx := bytes.LastIndex(data, paddingBlock)
				if slackPaddingIdx == -1 {
					fmt.Printf("\n%s ❌ Sieve burst: No global block open.%s\n", Red, Reset)
					continue
				}
				patchedData = append([]byte(nil), data...)
				copy(patchedData[slackPaddingIdx:], []byte(replaceStr))
				truncatedReplace := replaceStr[:len(findBytes)]
				copy(patchedData[idx:], []byte(truncatedReplace))
			} else {
				replaceBytes := []byte(replaceStr)
				paddingSize := availableContainerSize - len(replaceBytes)
				padding := make([]byte, paddingSize)
				replaceBytes = append(replaceBytes, padding...)
				targetPatchZone := data[idx : idx+availableContainerSize]
				patchedData = bytes.Replace(data, targetPatchZone, replaceBytes, 1)
			}
			
			fmt.Print("\033[H\033[2J")
			runLoadingBar("COMMITTING_CONFIRMED_DATA_BLOBS_TO_DISK")
			
			localOutputFile := targetFile + "_patched"
			err = os.WriteFile(localOutputFile, patchedData, 0755)
			if err != nil {
				fmt.Printf(" %s❌ Access Denied: File stream locked.%s\n", Red, Reset)
			} else {
				fmt.Printf(" %s⚡ VALIDATED TRANSACTIONS COMPLETED: %s%s\n", Green, localOutputFile, Reset)
				fmt.Printf(" %s📦 SNAPSHOT HOUSED AT: %s%s\n", Gray, backupPath, Reset)
			}
			fmt.Printf("\n%sPress [ENTER] to return...%s", Gray, Reset)
			_, _ = fmt.Scanln()
			
			
			
	case 4:
		clearTerminal()
		fmt.Printf("%s>>> AUTOMATED MATRIX DECOMPRESSION SYSTEM%s\n\n", Orange, Reset)
			
		data, err := os.ReadFile(targetFile)
		if err != nil {
			fmt.Printf(" %s❌ Access Malfunction: Unable to parse target bounds.%s\n", Red, Reset)
			fmt.Printf("\n%sPress [ENTER] to return...%s", Gray, Reset)
			_, _ = fmt.Scanln()
			continue
		}
			
		if !bytes.Contains(data, []byte("UPX!")) {
			fmt.Printf(" %s[-] Target Scan Clean: Standard UPX! headers are absent from this binary layout.%s\n", Yellow, Reset)
			fmt.Printf("    If file entropy calculation is spiked, it likely implements a custom cryptor.\n")
			fmt.Printf("\n%sPress [ENTER] to return...%s", Gray, Reset)
			_, _ = fmt.Scanln()
			continue
		}
			
		_, lookupErr := exec.LookPath("upx")
		if lookupErr != nil {
			fmt.Printf(" %s❌ COMPLIANCE BREAK: 'upx' decompression framework utility not found in system path.%s\n", Red, Reset)
			fmt.Printf("    To resolve this error on macOS, open a new shell terminal and execute: %sbrew install upx%s\n", Yellow, Reset)
			fmt.Printf("\n%sPress [ENTER] to return...%s", Gray, Reset)
			_, _ = fmt.Scanln()
			continue
		}
			
		fmt.Printf(" %s[+] Valid compressed header verified. Spawning rollback shell module...%s\n\n", Green, Reset)
		runLoadingBar("DEPLOYING_UPX_UNPACK_ROUTINES_AND_RESTORING_OFFSETS")
			
		cmd := exec.Command("upx", "-d", targetFile)
		var outBuf, errBuf bytes.Buffer
		cmd.Stdout = &outBuf
		cmd.Stderr = &errBuf
			
		execErr := cmd.Run()
			
		if execErr != nil {
			fmt.Printf(" %s❌ UNPACK_FAULT: Operating system compression reversal burst an exception error:%s\n", Red, Reset)
			fmt.Println(errBuf.String())
		} else {
			fmt.Printf(" %s✅ REGISTERS SYNCED: Binary decompression transaction successful!%s\n", Green, Reset)
			fmt.Printf("    Target payload is now fully unpacked and wide open for direct precision patches.\n")
			
			if updatedInfo, statErr := os.Stat(targetFile); statErr == nil {
				fileInfo = updatedInfo
			}
		}
			
		fmt.Printf("\n%sPress [ENTER] to return...%s", Gray, Reset)
		_, _ = fmt.Scanln()
			
	case 5:
		clearTerminal()
		fmt.Printf("%s>>> EXTRACTING UTF-16 WIDE CHARACTER STRING HARVEST%s\n\n", Cyan, Reset)
		runLoadingBar("SCANNING_UNICODE_PAYLOAD_CHANNELS")
			
		data, err := os.ReadFile(targetFile)
		if err != nil {
			fmt.Printf(" %s❌ Error reading workspace buffer.%s\n", Red, Reset)
		} else {
			// Internal local scanner logic to parse UTF-16 arrays matching printable boundaries
			var u16Strings []string
			var current []rune
			
			for i := 0; i < len(data)-1; i += 2 {
				// Read 2 bytes as little-endian uint16 unicode code point
				u16 := uint16(data[i]) | (uint16(data[i+1]) << 8)
				
				if u16 >= 0x20 && u16 <= 0x7E {
					current = append(current, rune(u16))
				} else if u16 == 0x00 && len(current) >= 4 {
					u16Strings = append(u16Strings, string(current))
					current = nil
				} else {
					current = nil
				}
			}
			
			if len(u16Strings) == 0 {
				fmt.Printf(" %s[-] Scan clean: Zero standard UTF-16 wide strings detected.%s\n", Yellow, Reset)
			} else {
				fmt.Printf(" %s[+] Discovered %d valid Unicode payloads:%s\n\n", Green, len(u16Strings), Reset)
				for _, str := range u16Strings {
					fmt.Printf("   » %s%s%s\n", White, str, Reset)
				}
			}
		}
		fmt.Printf("\n%sPress [ENTER] to return...%s", Gray, Reset)
		_, _ = fmt.Scanln()
			
			
		case 6:
			clearTerminal()
			fmt.Printf("%s>>> MACHINE OPCODE HEX INSTRUCTION MODIFICATION PANEL%s\n\n", Magenta, Reset)
			
			var findHex, replaceHex string
			fmt.Printf(" %sEnter Target Opcodes to Find (e.g. 740A):%s ", White, Reset)
			_, _ = fmt.Scanf("%s\n", &findHex)
			fmt.Printf(" %sEnter New Bypass Opcodes to Inject (e.g. 9090):%s ", White, Reset)
			_, _ = fmt.Scanf("%s\n", &replaceHex)
			
			findHex = strings.ReplaceAll(findHex, " ", "")
			replaceHex = strings.ReplaceAll(replaceHex, " ", "")
			
			if len(findHex) != len(replaceHex) {
				fmt.Printf("\n%s ❌ OFFSET BREACH: Hex arrays must maintain matching byte lengths to prevent shifting!%s\n", Red, Reset)
				fmt.Printf("\n%sPress [ENTER] to return...%s", Gray, Reset)
				_, _ = fmt.Scanln()
				continue
			}
			
			if len(findHex) == 0 || len(findHex)%2 != 0 {
				fmt.Printf("\n%s ❌ Error: Hex pairs must be complete and non-empty.%s\n", Red, Reset)
				fmt.Printf("\n%sPress [ENTER] to return...%s", Gray, Reset)
				_, _ = fmt.Scanln()
				continue
			}
			
			data, err := os.ReadFile(targetFile)
			if err != nil {
				fmt.Printf(" %s❌ Error reading binary target.%s\n", Red, Reset)
				continue
			}
			
			// Helper logic to safely parse hex strings into literal byte slices
			findBytes := make([]byte, len(findHex)/2)
			for i := 0; i < len(findBytes); i++ {
				val, _ := strconv.ParseUint(findHex[i*2:i*2+2], 16, 8)
				findBytes[i] = byte(val)
			}
			
			replaceBytes := make([]byte, len(replaceHex)/2)
			for i := 0; i < len(replaceBytes); i++ {
				val, _ := strconv.ParseUint(replaceHex[i*2:i*2+2], 16, 8)
				replaceBytes[i] = byte(val)
			}
			
			if !bytes.Contains(data, findBytes) {
				fmt.Printf("\n%s ❌ ERROR: Target opcode sequence absent from binary instructions layout.%s\n", Red, Reset)
				fmt.Printf("\n%sPress [ENTER] to return...%s", Gray, Reset)
				_, _ = fmt.Scanln()
				continue
			}
			
			backupPath, backupErr := generateBackup(targetFile)
			if backupErr != nil {
				fmt.Printf("\n%s ❌ Safety halt: Backup engine malfunction: %s%s\n", Red, backupErr.Error(), Reset)
				fmt.Printf("\n%sPress [ENTER] to return...%s", Gray, Reset)
				_, _ = fmt.Scanln()
				continue
			}
			
			runLoadingBar("MODIFYING_RAW_MACHINE_INSTRUCTION_OPCODES")
			
			patchedData := bytes.Replace(data, findBytes, replaceBytes, 1)
			localOutputFile := targetFile + "_patched"
			err = os.WriteFile(localOutputFile, patchedData, 0755)
			if err != nil {
				fmt.Printf(" %s❌ Access Denied: File stream locked.%s\n", Red, Reset)
			} else {
				fmt.Printf(" %s⚡ ASSEMBLY OPERATION MODIFIED SUCCESSFULLY: %s%s\n", Green, localOutputFile, Reset)
				fmt.Printf(" %s📦 BACKUP STORED AT: %s%s\n", Gray, backupPath, Reset)
			}
			fmt.Printf("\n%sPress [ENTER] to return...%s", Gray, Reset)
			_, _ = fmt.Scanln()
			
			
		case 7:
			clearTerminal()
			fmt.Printf("%s>>> MACH-O PRIVILEGE CODESIGN ENTITLEMENT SPLICER%s\n\n", Orange, Reset)
			
			if !strings.Contains(strings.ToLower(metadata), "mach-o") {
				fmt.Printf(" %s❌ HARDWARE FAULT: This operation requires an Apple Mach-O file container architecture.%s\n", Red, Reset)
				fmt.Printf("\n%sPress [ENTER] to return...%s", Gray, Reset)
				_, _ = fmt.Scanln()
				continue
			}
			
			fmt.Printf(" [*] Harvesting native embedded XML verification blueprints from '%s'...\n", targetFile)
			xmlCmd := exec.Command("codesign", "-d", "--entitlements", "extracted_ents.xml", targetFile)
			_ = xmlCmd.Run()
			
			runLoadingBar("SPLICING_XML_PRIVILEGES_AND_AD_HOC_SIGNING")
			
			var signCmd *exec.Command
			localOutputFile := targetFile + "_patched"
			
			if _, err := os.Stat("extracted_ents.xml"); err == nil {
				fmt.Printf(" %s[+] Intact XML entitlements recovered. Resigning with full privilege matrix templates...%s\n", Green, Reset)
				signCmd = exec.Command("codesign", "--force", "--sign", "-", "--entitlements", "extracted_ents.xml", localOutputFile)
				_ = os.Remove("extracted_ents.xml")
			} else {
				fmt.Printf(" %s[i] Generic layout signature detected. Enforcing custom fallback ad-hoc seal...%s\n", Yellow, Reset)
				signCmd = exec.Command("codesign", "--force", "--sign", "-", localOutputFile)
			}
			
			err = signCmd.Run()
			if err != nil {
				fmt.Printf(" %s❌ SIGNING_FAULT: Operating system cryptographic alignment burst an exception error.%s\n", Red, Reset)
			} else {
				fmt.Printf(" %s🛡️  HARDENING VALIDATED: Security signature seal refreshed successfully on targeted file!%s\n", Green, Reset)
			}
			fmt.Printf("\n%sPress [ENTER] to return...%s", Gray, Reset)
			_, _ = fmt.Scanln()
			
		case 8:
			clearTerminal()
			fmt.Printf("%s [!] Terminal session safely detached. Goodbye Operator.%s\n\n", Red, Reset)
			os.Exit(0)
			
		default:
			fmt.Printf("%sInvalid choice vector.%s\n", Red, Reset)
			time.Sleep(1 * time.Second)
		}
	}
}



			
			