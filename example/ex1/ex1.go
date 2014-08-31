package main

import (
	"fmt"
	"gopkg.in/wu.v0"
	"strings"
)

func main() {
	ses := wu.NewSession()
	sea := ses.NewSearcher(true)
	updates := sea.Query(wu.NewSearchQuery().IsInstalled(true).IsHidden(false).String())
	fmt.Printf("Updates: %d\n", len(updates.Updates))
	for _, update := range updates.Updates {
		fmt.Printf("[%s]\n", update.UpdateID)
		fmt.Printf("[%s] %s\n", update.Type, update.Title)
		fmt.Printf("[%s]\n", update.UpdateID)
		if update.MsrcSeverity != "" {
			fmt.Printf("\tMsrcSeverity: %s\n", update.MsrcSeverity)
		}
		if !update.Deadline.IsZero() {
			fmt.Printf("\tDeadline: %s\n", update.Deadline)
		}
		if !update.LastDeploymentChangeTime.IsZero() {
			fmt.Printf("\tLastDeploymentChangeTime: %s\n", update.LastDeploymentChangeTime)
		}
		fmt.Printf("\tAutoDownload: %s\n", update.AutoDownload)
		fmt.Printf("\tAutoSelection: %s\n", update.AutoSelection)
		fmt.Printf("\tIsInstalled: %t\n", update.IsInstalled)
		fmt.Printf("\tIsUninstallable: %t\n", update.IsUninstallable)
		fmt.Printf("\tIsHidden: %t\n", update.IsHidden)
		fmt.Printf("\tIsPresent: %t\n", update.IsPresent)
		fmt.Printf("\tEulaAccepted: %t\n", update.EulaAccepted)
		fmt.Printf("\tRebootRequired: %t\n", update.RebootRequired)
		fmt.Printf("\tMaxDownloadSize: %d bytes\n", update.MaxDownloadSize)
		fmt.Printf("\tMinDownloadSize: %d bytes\n", update.MinDownloadSize)
		if len(update.SupersededUpdateIDs) > 0 {
			fmt.Printf("\tSupersededUpdateIDs:\n\t\t%s\n", strings.Join(update.SupersededUpdateIDs, "\n\t\t"))
		}
		if len(update.SecurityBulletinIDs) > 0 {
			fmt.Printf("\tSecurityBulletinIDs:\n\t\t%s\n", strings.Join(update.SecurityBulletinIDs, "\n\t\t"))
		}
		if len(update.KBArticleIDs) > 0 {
			fmt.Printf("\tKBArticleIDs:\n\t\t%s\n", strings.Join(update.KBArticleIDs, "\n\t\t"))
		}
		if len(update.CveIDs) > 0 {
			fmt.Printf("\tCveIDs:\n\t\t%s\n", strings.Join(update.CveIDs, "\n\t\t"))
		}
		if len(update.Categories.Categories) > 0 {
			fmt.Print("\tCategories:\n")
			for _, category := range update.Categories.Categories {
				fmt.Printf("\t\t[%s] %s\n", category.CategoryID, category.Name)
				//fmt.Printf("\t\t\t%s\n", category.Description)
			}
		}
		fmt.Printf("\tRecommendedCPUSpeed: %d MHz\n", update.RecommendedCPUSpeed)
		fmt.Printf("\tRecommendedHardDiskSpace: %d MB\n", update.RecommendedHardDiskSpace)
		fmt.Printf("\tRecommendedMemory: %d MB\n", update.RecommendedMemory)
		fmt.Println("\n")
	}
}
