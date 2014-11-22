package models

import (
    "strconv"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const COLLECTION = "projects"


type Team struct {
    Id       bson.ObjectId `bson:"_id,omitempty"`
    Name     string        `bson:"name"`
    Users    []string      `bson:"users"`
    Projects []string      `bson:"projects"`
}

type Project struct {
    Id      bson.ObjectId `bson:"_id,omitempty"`
    Name    string        `bson:"name"`
    Users   mgo.DBRef     `bson:"users"`
    Todos   []mgo.DBRef      `bson:"todos"`
}

type Todo struct {
    Id        bson.ObjectId `bson:"_id,omitempty"`
    AuthorId  string        `bson:"author_id"`
    Title     string        `bson:"title"`
    CDate     string        `bson:"c_date"`
    SSDate    string        `bson:"s_s_date"`
    SFDate    string        `bson:"s_f_date"`
    FDate     string        `bson:"f_date"`
    Finished  bool          `bson:"finished"`
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
func FindTeamList(d *mgo.Database) []TeamListResponse {
    teamList := []TeamList{}
    result   := []TeamListResponse{}
    Collection(d).Find(nil).All(&teamList)

    // TeamList to TeamListResponse
    for _, each := range teamList {
        newTeamList := TeamListResponse{}
        newTeamList.Id = each.Id.Hex()
        newTeamList.Name = each.Name
        result = append(result, newTeamList)
    }

    return result
}

// get project list (API_P001)
func FindProjectListByTeamId(d *mgo.Database, HexTeamId string) []ProjectListResponse {
    projectList := []ProjectList{}
    result      := []ProjectListResponse{}
    if bson.IsObjectIdHex(HexTeamId) {
        Id := bson.ObjectIdHex(HexTeamId)
        Collection(d).FindId(Id).All(&projectList)
    }

    for _, each := range projectList {
        newPList := ProjectListResponse{}
        newPList.Id = each.Id.Hex()
        newPList.Name = each.Name
        result = append(result, newPList)
    }
    return result
}


// get todo list (API_T001)
func FindTodoListByProjectId(d *mgo.Database, projectId string) []TodoListResponse {
    todoList := []TodoListResponse{}

    if bson.IsObjectIdHex(projectId) {
        project := Project{}
        id := bson.ObjectIdHex(projectId)
        Collection(d).FindId(id).One(&project)

        for _, todoRef := range project.Todos {
            todo := Todo{}
            FindRef(d, &todoRef).One(&todo)

            nTodoRes := TodoListResponse{}
            nTodoRes.Id = bson.ObjectId.Hex(todo.Id)
            nTodoRes.AuthorName = todo.AuthorId
            nTodoRes.CDate = todo.CDate
            nTodoRes.SSDate = todo.SSDate
            nTodoRes.SFDate = todo.SFDate
            nTodoRes.FDate = todo.FDate
            nTodoRes.Finished = strconv.FormatBool(todo.Finished)

            todoList = append(todoList, nTodoRes)
        }
    }

    return todoList
}

