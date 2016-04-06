package core

import (
  "github.com/dev-op-spec/engine/core/ports"
  "github.com/dev-op-spec/engine/core/models"
)

type setDescriptionOfOpUseCase interface {
  Execute(
  req models.SetDescriptionOfOpReq,
  ) (err error)
}

func newSetDescriptionOfOpUseCase(
filesys ports.Filesys,
pathToOpFileFactory pathToOpFileFactory,
yamlCodec yamlCodec,
) setDescriptionOfOpUseCase {

  return &_setDescriptionOfOpUseCase{
    filesys:filesys,
    pathToOpFileFactory:pathToOpFileFactory,
    yamlCodec:yamlCodec,
  }

}

type _setDescriptionOfOpUseCase struct {
  filesys                  ports.Filesys
  pathToOpFileFactory pathToOpFileFactory
  yamlCodec                yamlCodec
}

func (this _setDescriptionOfOpUseCase) Execute(
req models.SetDescriptionOfOpReq,
) (err error) {

  pathToOpFile := this.pathToOpFileFactory.Construct(
    req.ProjectUrl,
    req.OpName,
  )

  opFileBytes, err := this.filesys.GetBytesOfFile(pathToOpFile)
  if (nil != err) {
    return
  }

  opFile := opFile{}
  err = this.yamlCodec.fromYaml(
    opFileBytes,
    &opFile,
  )
  if (nil != err) {
    return
  }

  opFile.Description = req.Description

  opFileBytes, err = this.yamlCodec.toYaml(&opFile)
  if (nil != err) {
    return
  }

  err = this.filesys.SaveFile(
    pathToOpFile,
    opFileBytes,
  )

  return

}
