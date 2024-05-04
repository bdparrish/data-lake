package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bufbuild/protovalidate-go"
	"github.com/codingexplorations/data-lake/common/pkg/converter"
	"github.com/codingexplorations/data-lake/common/pkg/log"
	"github.com/codingexplorations/data-lake/common/pkg/models/v1/db"
	"github.com/codingexplorations/data-lake/common/pkg/models/v1/proto"
	"github.com/codingexplorations/data-lake/object-service/pkg/repository"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

type ObjectController struct {
	repository      repository.ObjectRepository[db.Object]
	logger          log.Logger
	marshaller      protojson.MarshalOptions
	protoValidator  *protovalidate.Validator
	objectConverter converter.GenericConverter[proto.Object, db.Object]
}

func NewObjectController(
	logger log.Logger,
	repository repository.ObjectRepository[db.Object],
) *ObjectController {
	validator, err := protovalidate.New()
	if err != nil {
		logger.Error(fmt.Sprintf("error in creating proto validator: %v", err))
		return nil
	}

	controller := ObjectController{
		repository: repository,
		logger:     logger,
		marshaller: protojson.MarshalOptions{
			EmitUnpopulated: true,
		},
		protoValidator:  validator,
		objectConverter: converter.JsonMarshallingConverter[proto.Object, db.Object]{},
	}
	return &controller
}

// GetObjects godoc
//
//	@Summary		Get objects in the system. Utilize page and size to paginate through the list of objects. Page and size are optional as the defaults for the system will be used.
//	@Description	Queries the API for objects and details.
//	@Tags			object
//	@Accept			json
//	@Produce		json
//	@Param			page	query	int	false	"Page"
//	@Param			size	query	int	false	"Page Size"
//	@Success		200		object	proto.Object
//	@Router			/ [get]
func (controller *ObjectController) GetObjects(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		controller.logger.Error(fmt.Sprintf("error in parsing page: %v", err))
		return
	}
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil {
		controller.logger.Error(fmt.Sprintf("error in parsing size: %v", err))
		return
	}

	objects, err := controller.repository.GetAll(page, size)
	if err != nil {
		controller.logger.Error(fmt.Sprintf("error in getting objects: %v", err))
		return
	}

	if len(objects) == 0 {
		c.Data(http.StatusOK, "application/json", []byte{})
		return
	}

	protoObjects, err := controller.objectConverter.DbToProtoSlice(objects)
	if err != nil {
		controller.logger.Error(fmt.Sprintf("error in converting db objects to proto objects: %v", err))
	}

	response := proto.ObjectGetAllResponse{
		Objects: protoObjects,
	}

	payload, _ := controller.marshaller.Marshal(&response)
	c.Data(http.StatusOK, "application/json", payload)
}
