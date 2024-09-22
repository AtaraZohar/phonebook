package services

import (
	"phonebook/models"

	"gorm.io/gorm"
)

type ContactService struct {
	DB *gorm.DB
}

func NewContactService(db *gorm.DB) *ContactService {
	return &ContactService{DB: db}
}

func (s *ContactService) CreateContact(contact *models.Contact) error {
	return s.DB.Create(contact).Error
}

func (s *ContactService) GetContacts(page, limit int) ([]models.Contact, error) {
	var contacts []models.Contact
	offset := (page - 1) * limit

	err := s.DB.Limit(limit).Offset(offset).Find(&contacts).Error
	if err != nil {
		return nil, err
	}

	return contacts, nil
}

func (s *ContactService) GetContactByID(id uint) (*models.Contact, error) {
	var contact models.Contact
	err := s.DB.First(&contact, id).Error
	return &contact, err
}

func (s *ContactService) SearchContacts(firstName, lastName, phoneNumber, address string) ([]models.Contact, error) {
	var contacts []models.Contact
	query := s.DB

	if firstName != "" {
		query = query.Where("first_name ILIKE ?", "%"+firstName+"%")
	}
	if lastName != "" {
		query = query.Where("last_name ILIKE ?", "%"+lastName+"%")
	}
	if phoneNumber != "" {
		query = query.Where("phone_number ILIKE ?", "%"+phoneNumber+"%")
	}
	if address != "" {
		query = query.Where("address ILIKE ?", "%"+address+"%")
	}

	err := query.Find(&contacts).Error
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

func (s *ContactService) UpdateContact(id uint, contact models.Contact) (models.Contact, error) {
	var existingContact models.Contact
	if err := s.DB.First(&existingContact, id).Error; err != nil {
		return models.Contact{}, err
	}

	existingContact.FirstName = contact.FirstName
	existingContact.LastName = contact.LastName
	existingContact.PhoneNumber = contact.PhoneNumber
	existingContact.Address = contact.Address

	if err := s.DB.Save(&existingContact).Error; err != nil {
		return models.Contact{}, err
	}

	return existingContact, nil
}

func (s *ContactService) DeleteContact(id uint) error {
	return s.DB.Delete(&models.Contact{}, id).Error
}
