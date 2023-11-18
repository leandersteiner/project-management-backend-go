package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Login(ctx echo.Context) error {
	return nil
}

func (h *Handler) GetAllOrganisations(ctx echo.Context) error {
	return nil
}

func (h *Handler) CreateOrganisation(ctx echo.Context) error {
	return nil
}

func (h *Handler) DeleteOrganisation(ctx echo.Context, orgId string) error {
	return nil
}

func (h *Handler) GetOrganisation(ctx echo.Context, orgId string) error {
	return nil
}

func (h *Handler) AddUserToOrganisation(ctx echo.Context, orgId string) error {
	return nil
}

func (h *Handler) GetAllProjectsInOrganisation(ctx echo.Context, orgId string) error {
	return nil
}

func (h *Handler) GetAllTeamsInOrg(ctx echo.Context, orgId string) error {
	return nil
}

func (h *Handler) CreateTeam(ctx echo.Context, orgId string) error {
	return nil
}

func (h *Handler) DeleteTeam(ctx echo.Context, orgId string, teamId string) error {
	return nil
}

func (h *Handler) GetTeam(ctx echo.Context, orgId string, teamId string) error {
	return nil
}

func (h *Handler) AddUserToTeam(ctx echo.Context, orgId string, teamId string) error {
	return nil
}

func (h *Handler) GetAllProjectsInTeam(ctx echo.Context, orgId string, teamId string) error {
	return nil
}

func (h *Handler) CreateProject(ctx echo.Context, orgId string, teamId string) error {
	return nil
}

func (h *Handler) DeleteProject(ctx echo.Context, orgId string, teamId string, projectId string) error {
	return nil
}

func (h *Handler) AddUserToProject(ctx echo.Context, orgId string, teamId string, projectId string) error {
	return nil
}

func (h *Handler) GetAllProjectsForUser(ctx echo.Context) error {
	return nil
}

func (h *Handler) GetProject(ctx echo.Context, projectId string) error {
	return nil
}

func (h *Handler) GetBacklog(ctx echo.Context, projectId string) error {
	return nil
}

func (h *Handler) GetBoard(ctx echo.Context, projectId string) error {
	return nil
}

func (h *Handler) GetAllColumns(ctx echo.Context, projectId string) error {
	return nil
}

func (h *Handler) CreateColumn(ctx echo.Context, projectId string) error {
	return nil
}

func (h *Handler) DeleteColumn(ctx echo.Context, projectId string, columnId string) error {
	return nil
}

func (h *Handler) GetColumn(ctx echo.Context, projectId string, columnId string) error {
	return nil
}

func (h *Handler) UpdateColumn(ctx echo.Context, projectId string, columnId string) error {
	return nil
}

func (h *Handler) GetAllTasksInColumn(ctx echo.Context, projectId string, columnId string) error {
	return nil
}

func (h *Handler) CreateTaskInColumn(ctx echo.Context, projectId string, columnId string) error {
	return nil
}

func (h *Handler) GetAllProjectMembers(ctx echo.Context, projectId string) error {
	return nil
}

func (h *Handler) GetAllSprints(ctx echo.Context, projectId string) error {
	return nil
}

func (h *Handler) CreateSprint(ctx echo.Context, projectId string) error {
	return nil
}

func (h *Handler) DeleteSprint(ctx echo.Context, projectId string, sprintId string) error {
	return nil
}

func (h *Handler) GetSprint(ctx echo.Context, projectId string, sprintId string) error {
	return nil
}

func (h *Handler) UpdateSprint(ctx echo.Context, projectId string, sprintId string) error {
	return nil
}

func (h *Handler) GetAllTasks(ctx echo.Context, projectId string) error {
	return nil
}

func (h *Handler) CreateTask(ctx echo.Context, projectId string) error {
	return nil
}

func (h *Handler) CreateTaskState(ctx echo.Context, projectId string) error {
	return nil
}

func (h *Handler) DeleteTask(ctx echo.Context, projectId string, taskId string) error {
	return nil
}

func (h *Handler) GetTask(ctx echo.Context, projectId string, taskId string) error {
	return nil
}

func (h *Handler) UpdateTask(ctx echo.Context, projectId string, taskId string) error {
	return nil
}

func (h *Handler) Register(ctx echo.Context) error {
	return nil
}

func (h *Handler) CreateTaskComment(ctx echo.Context, taskId string) error {
	return nil
}

func (h *Handler) DeleteTaskComment(ctx echo.Context, taskId string, commentId string) error {
	return nil
}

func (h *Handler) UpdateTaskComment(ctx echo.Context, taskId string, commentId string) error {
	return nil
}

func (h *Handler) CreateSubtask(ctx echo.Context, taskId string) error {
	return nil
}

func (h *Handler) DeleteSubtask(ctx echo.Context, taskId string, subtaskId string) error {
	return nil
}

func (h *Handler) UpdateSubtask(ctx echo.Context, taskId string, subtaskId string) error {
	return nil
}

func (h *Handler) GetAllTeams(ctx echo.Context) error {
	return nil
}

func (h *Handler) GetAllUsers(ctx echo.Context) error {
	ctx.JSON(http.StatusAccepted, "hi")
	return nil
}

func (h *Handler) DeleteUser(ctx echo.Context, userId string) error {
	return nil
}

func (h *Handler) GetUser(ctx echo.Context, userId string) error {
	return nil
}

func (h *Handler) UpdateUser(ctx echo.Context, userId string) error {
	return nil
}
