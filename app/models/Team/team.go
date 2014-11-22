package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const COLLECTION = "teams"


type Team struct {
    Id       bson.ObjectId `bson:"_id,omitempty"`
    Name     string        `bson:"name"`
    Users    []string      `bson:"users"`
    Projects []mgo.DBRef     `bson:"projects"`
}

type Project struct {
    Id      bson.ObjectId `bson:"_id,omitempty"`
    Name    string        `bson:"name"`
    Users   []mgo.DBRef   `bson:"users"`
    Todos   []string      `bson:"todos"`
}

type Todo struct {
    Id        bson.ObjectId `bson:"_id,omitempty"`
    AuthorId  string        `bson:"author_id"`
    Title     string        `bson:"title"`
    CDate     string        `bson:"c_date"`
    SSDate    string        `bson:"s_s_date"`
    SFDate    string        `bson:"s_f_date"`
    FDate     string        `bson:"f_date"`
    Finished  string        `bson:"finished"`
}












type TeamList struct {
    Id    bson.ObjectId  `bson:"_id,omitempty"`
	Name  string         `bson:"name"`
}

type ProjectList struct {
    Id        bson.ObjectId      `bson:"_id,omitempty"`
    Name      string             `bson:"name"`
	Users     []string           `bson:"users"`
	Projects  []ProjectListItem  `bson:"projects"`
}

type ProjectListItem struct {
    Id    string  `bson:"id"`
    Name  string  `bson:"name"`
}



type TeamListResponse struct {
    Id    string
    Name  string
}

type ProjectListResponse struct {
    Id    string
    Name  string
}

type TodoListResponse struct {
    Id          string
    AuthorName  string
    Title       string 
    CDate       string
    SSDate      string
    SFDate      string
    FDate       string
    Finished    string
}

func FindRef(d *mgo.Database, ref *mgo.DBRef) *mgo.Query {
	var c *mgo.Collection

	if ref.Database == "" {
		c = d.C(ref.Collection)
	} else {
		c = d.Session.DB(ref.Database).C(ref.Collection)
	}

    id := bson.ObjectIdHex(ref.Id.(string))
    return c.FindId(id)
}

func Collection(d *mgo.Database) *mgo.Collection {
	return d.C(COLLECTION)
}

// get team list (API_G001)
func FindTeamList(d *mgo.Database) []Team {
    teamList := []Team{}
    result   := []TeamListResponse{}
    Collection(d).Find(nil).All(&teamList)

    // TeamList to TeamListResponse
    for _, each := range teamList {
        newTeamList := TeamListResponse{}
        newTeamList.Id = each.Id.Hex()
        newTeamList.Name = each.Name
        result = append(result, newTeamList)
    }


    return teamList
}

// get project list (API_P001)
func FindProjectListByTeamId(d *mgo.Database, HexTeamId string) []ProjectListResponse {
    team := Team{}
    projectList := []ProjectListResponse{}

    //result      := []ProjectListResponse{}
    if bson.IsObjectIdHex(HexTeamId) {
        Id := bson.ObjectIdHex(HexTeamId)
        Collection(d).FindId(Id).One(&team)
    }

    for _, projectRef := range team.Projects {
        project := Project{}
        FindRef(d, &projectRef).One(&project)

        nProjectRef := ProjectListResponse{}
        nProjectRef.Id = bson.ObjectId.Hex(project.Id)  // convert OID to Hex
        nProjectRef.Name = project.Name

        projectList = append(projectList, nProjectRef)
    }

    //pro := Project{}
    //FindRef(d, &teamList[0].Projects[0]).One(&pro)

    return projectList
}

// get todo list (API_T001)
func FindTodoListByTeamAndProjectId(d *mgo.Database, TeamHexId string, ProjectId string) []TodoListResponse {
    result := []TodoListResponse{}
//    pResult := Team{}
//    if bson.IsObjectIdHex(TeamHexId) {
//        Id := bson.ObjectIdHex(TeamHexId)
//        Collection(d).FindId(Id).One(&pResult)
//
//        for _, Project := range pResult.Projects {
//            if Project.Id == ProjectId {
//                todo := TodoListResponse{}
//
//                for _, Todo := range Project.Todos {
//                    todo.Id = Todo.Id.Hex()
//                    todo.AuthorName = "テスト"
//                    todo.Title      = Todo.Title
//                    todo.CDate      = Todo.CDate
//                    todo.SSDate     = Todo.SSDate
//                    todo.SFDate     = Todo.SFDate
//                    todo.FDate      = Todo.FDate
//                    todo.Finished   = Todo.Finished
//                }
//
//                result = append(result, todo)
//
//            }
//        }
//
//    }
    return result
}

