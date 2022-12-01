export interface StatisticsData {
  design_project_count: number
  libraries_count: number
  design_results_count: number
  user_count: number
  design_project_delivered_count: number
  design_project_exceeded_count: number
  design_project_imported_count: number
  design_project_created_count: number
  design_project_started_count: number
  design_project_finished_count: number
  design_project_review_unaccepted_count: number
  design_project_reviewed_count: number
  design_project_approve_unaccepted_count: number
  design_project_approved_count: number
  design_project_check_unaccepted_count: number
  design_project_checked_count: number
  design_project_overdue_count: number
  design_project_weekly_status_count: Array<{
    date: string
    design_project_count: number
  }>
}
