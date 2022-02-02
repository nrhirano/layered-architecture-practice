package factory

import (
	"github.com/go-xorm/xorm"
	"github.com/nso-rnt/td-equ-api/registry"
	"github.com/nso-rnt/td-equ-api/service"
)

const (
	// ServiceKey はサービスファクトリ取得キー名
	ServiceKey = "service_factory"
)

// Servicer はサービスファクトリ
type Servicer interface {
	NewUsers() service.UsersInterface
	NewEquipments() service.EquipmentsInterface
	NewEquipmentCode() service.EquipmentCodeInterface
	NewInspectionTemplates() service.InspectionTemplatesInterface
	NewGeneralAttributeItems() service.GeneralAttributeItemsInterface
	NewGeneralAttributes() service.GeneralAttributesInterface
	NewInspectionsTG() service.InspectionsTGInterface
	NewSites() service.SitesInterface
	NewInspectionAlert() service.InspectionAlertInterface
	NewAccessories() service.AccessoriesInterface
	NewReservations() service.ReservationsInterface
	NewAccessoryGroups() service.AccessoryGroupsInterface
	NewTableLayouts() service.TableLayoutsInterface
	NewEquipmentGr() service.EquipmentGrInterface
	NewEquipmentCodeFiles() service.EquipmentCodeFilesInterface
	NewEquipmentFiles() service.EquipmentFilesInterface
	NewFiles() service.FilesInterface
	NewReports() service.ReportsInterface
	NewUserLocks() service.UserLocksInterface
	NewEquipmentHistories() service.EquipmentHistoriesInterface
	NewLeavingPicking() service.LeavingPickingsInterface
	NewLeavings() service.LeavingsInterface
	NewAntiRusts() service.AntiRustsInterface
	NewEnteringPickings() service.EnteringPickingsInterface
	NewEnterings() service.EnteringsInterface
	NewProjects() service.ProjectsInterface
	NewBatchHistories() service.BatchHistoriesInterface
	NewInspections() service.InspectionsInterface
	NewEnteringUncheckedItems() service.EnteringUncheckedItemsInterface
	NewAdditionalAccessories() service.AdditionalAccessoriesInterface
	NewRepairOutside() service.RepairOutsideInterface
	NewRepairOutsideN() service.RepairOutsideNInterface
	NewLendings() service.LendingsInterface
	NewBatch() service.BatchInterface
	NewLoadchain() service.LoadchainInterface
	NewUncontrolledItems() service.UncontrolledItemsInterface
	NewControls() service.ControlsInterface
	NewInventories() service.InventoriesInterface
}

// ServiceFactorySettings はサービスファクトリの設定
// 全ての設定が必須
type ServiceFactorySettings struct {
	Engine xorm.EngineInterface
}

// ServiceFactory はサービスファクトリの実装
// インフラ層の依存情報を初期化時に注入する
type ServiceFactory struct {
	settings  *ServiceFactorySettings
	repoMaker registry.RepositoryMaker
}

// NewService initializes factory with injected infra.
func NewService(settings *ServiceFactorySettings) *ServiceFactory {
	r := &ServiceFactory{
		settings: settings,
	}

	r.repoMaker = registry.NewRepository(&registry.RepositorySettings{Engine: settings.Engine})
	return r
}
